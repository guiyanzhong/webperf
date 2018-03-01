require 'sinatra'

set :environment, :production

get '/' do
  'hello world'
end
