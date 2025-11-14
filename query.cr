require "http/client"

class Query
  BASE_URL = "https://api.rambase.net"

  getter url : String
  getter token : String

  def initialize(endpoint : String, token : String)
    # Odcinamy baseUrl jeśli user podał pełny adres
    end_url = endpoint.sub(BASE_URL, "")
    @url = "#{BASE_URL}/#{end_url}".gsub("//", "/")

    @token = token
  end

  private def execute_fetch(body : Hash(String, String | JSON::Any)?)
    headers = HTTP::Headers{
      "Authorization" => "Bearer #{@token}",
      "Content-Type"  => "application/json",
    }

    response =
      if body
        method = body["method"].as_s || "POST"
        raw_body = body["body"]?
        HTTP::Client.post(@url, headers: headers, body: raw_body.try(&.to_s))
      else
        HTTP::Client.get(@url, headers: headers)
      end

    if response.status_code == 401
      AuthContext.set_auth("isLoggedIn", false)
      raise "Unauthorized"
    end

    JSON.parse(response.body)
  end

  def get_list
    execute_fetch(nil)
  end

  def post(request : JSON::Any?)
    body = if request
             {
               "method" => "POST",
               "body"   => request.to_json,
             }
           else
             {"method" => "POST"}
           end

    execute_fetch(body)
  end

  def put(request : JSON::Any?)
    body = if request
             {
               "method" => "PUT",
               "body"   => request.to_json,
             }
           else
             {"method" => "PUT"}
           end

    execute_fetch(body)
  end
end
