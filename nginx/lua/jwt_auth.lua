local jwt = require("resty.jwt")
local cjson = require("cjson")

local _M = {}

local secret = os.getenv("JWT_SECRET") or "default-secret-please-change"

function _M.verify()
    -- Skip JWT verification for specific paths
    print(secret)
    local skip_paths = {
        ["^/grafana"] = true,
        ["^/weaviate"] = true,
        ["^/.well%-known/acme%-challenge"] = true
    }
    
    for path, _ in pairs(skip_paths) do
        if ngx.var.uri:match(path) then
            return true
        end
    end

    local auth_header = ngx.var.http_Authorization
    if not auth_header then
        ngx.status = ngx.HTTP_UNAUTHORIZED
        ngx.header["Content-Type"] = "application/json"
        ngx.say(cjson.encode({error = "Missing Authorization header"}))
        return ngx.exit(ngx.HTTP_UNAUTHORIZED)
    end

    local _, _, token = string.find(auth_header, "Bearer%s+(.+)")
    if not token then
        ngx.status = ngx.HTTP_UNAUTHORIZED
        ngx.header["Content-Type"] = "application/json"
        ngx.say(cjson.encode({error = "Malformed Authorization header"}))
        return ngx.exit(ngx.HTTP_UNAUTHORIZED)
    end

    local jwt_obj = jwt:verify(secret, token)
    if not jwt_obj.verified then
        ngx.status = ngx.HTTP_UNAUTHORIZED
        ngx.header["Content-Type"] = "application/json"
        ngx.say(cjson.encode({error = "Invalid token: " .. (jwt_obj.reason or "unknown")}))
        return ngx.exit(ngx.HTTP_UNAUTHORIZED)
    end

    ngx.ctx.user = jwt_obj.payload
    return true
end

return _M