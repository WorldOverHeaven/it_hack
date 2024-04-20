package ru.mephi.auth.resource;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.core.Response;
import org.springframework.stereotype.Component;

@Path(value = "/auth")
@Component
public class AuthResource {

  @POST
  @Path("/create_user")
  public Response registerNewAccount(String login) {
    return null;
  }

  @POST
  @Path("/create_user")
  public Response registerCloud(String login) {
    return null;
  }

  @POST
  @Path("/create_user")
  public Response loginCloud(String login, String pass) {
    return null;
  }



  @GET
  @Path("/get_token")
  public Response getToken(String login) {
    return null;
  }





}
