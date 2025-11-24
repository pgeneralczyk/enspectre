require "http/client"

class Query
  BASE_URL = "https://api.rambase.net"

  getter url : String
  getter token : String

  def initialize(endpoint : String, token : String)
    # Odcinamy baseUrl jeśli user podał pełny adres
    end_url = endpoint.sub(BASE_URL, "")
    @url = "#{BASE_URL}/#{end_url}"
    @token = token
  end

  private def execute_fetch(body : Hash(String, String | JSON::Any)?)
    headers = HTTP::Headers{
      "Authorization" => "Bearer #{@token}",
      "Content-Type"  => "application/json",
    }

    response =
      if body
        method = body["method"]? || "POST"
        raw_body = body["body"]?
        HTTP::Client.post(@url, headers: headers, body: raw_body.try(&.to_s))
      else
        HTTP::Client.get(@url, headers: headers)
      end

    if response.status_code == 401
      raise "Unauthorized"
    end

    JSON.parse(response.body)
  end

  def get_list
    execute_fetch(nil)
  end

  def post(request : JSON::Any?)
    headers = HTTP::Headers{
      "Content-Type"  => "application/json",
      "Authorization" => "Bearer #{token}",
    }
    response = HTTP::Client.post(
      url,
      headers: headers,
      body: request.to_json
    )
  end

  def put(request : JSON::Any?)
    headers = HTTP::Headers{
      "Content-Type"  => "application/json",
      "Authorization" => "Bearer #{token}",
    }
    response = HTTP::Client.put(
      url,
      headers: headers,
      body: request.to_json
    )
  end

end
