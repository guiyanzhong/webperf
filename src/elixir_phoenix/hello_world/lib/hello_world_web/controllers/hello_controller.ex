defmodule HelloWorldWeb.HelloController do
  use HelloWorldWeb, :controller

  def hello_world(conn, _params) do
    text conn, "hello world"
  end

end
