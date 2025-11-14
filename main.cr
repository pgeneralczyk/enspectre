require "./query"

def use_get_list(endpoint : String, token : String)
  Query.new(endpoint,token).get_list
end

def use_post(endpoint : String, token : String, request)
  Query.new(endpoint,token).post(request)
end

def use_put(endpoint : String, token : String, request)
  Query.new(endpoint,token).put(request)
end


# teraz możesz używać Query:
puts "Hello World"

