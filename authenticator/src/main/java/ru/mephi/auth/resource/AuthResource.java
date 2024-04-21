package ru.mephi.auth.resource;

import jakarta.inject.Inject;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import org.springframework.stereotype.Component;
import ru.mephi.auth.service.AuthService;

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
    authService.joinCloud(cloudLogin, cloudPass);
    return Response.noContent().build();
  }

  @POST
  @Path(value = "/register_new_user")
  public Response registerNewUser(@QueryParam("login") String login) throws Exception {
    authService.registerNewUser(login);
    return Response.noContent().build();
  }

  @POST
  @Path(value = "/auth_user")
  public Response authUser(@QueryParam("login") String login) throws Exception {
    authService.authUser(login);
    return Response.noContent().build();
  }

  @POST
  @Path(value = "/verify_auth")
  @Produces(MediaType.APPLICATION_JSON)
  public Response verify(@QueryParam("login") String login) {
    return Response.ok(authService.verify(login)).build();
  }
}
