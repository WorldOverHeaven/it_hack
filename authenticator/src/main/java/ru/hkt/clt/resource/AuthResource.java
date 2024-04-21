package ru.hkt.clt.resource;

import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import org.springframework.stereotype.Component;
import ru.hkt.clt.service.AuthService;

@Component
@Path(value = "/auth")
public class AuthResource {
  private final AuthService authService;

  public AuthResource(AuthService authService) {
    this.authService = authService;
  }

  @POST
  @Path(value = "/join_cloud")
  public Response joinCloud(@QueryParam("cloudLogin") String cloudLogin, @QueryParam("cloudPass") String cloudPass) {
    return Response.ok(authService.joinCloud(cloudLogin, cloudPass)).build();
  }

  @POST
  @Path(value = "/register_new_user")
  public Response registerNewUser(@QueryParam("login") String login) throws Exception {
    return Response.ok(authService.registerNewUser(login)).build();
  }

  @POST
  @Path(value = "/auth_user")
  public Response authUser(@QueryParam("login") String login) throws Exception {
    return Response.ok(authService.authUser(login)).build();
  }

  @POST
  @Path(value = "/verify_auth")
  @Produces(MediaType.APPLICATION_JSON)
  public Response verify(@QueryParam("login") String login) {
    return Response.ok(authService.verify(login)).build();
  }
}