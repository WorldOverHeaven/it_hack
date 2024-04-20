package ru.mephi.auth.resource;

import jakarta.ws.rs.Consumes;
import jakarta.ws.rs.DELETE;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.PUT;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import ru.mephi.auth.dto.TaskDto;
import ru.mephi.auth.exception.IllegalTaskException;
import ru.mephi.auth.service.TaskService;

import java.util.List;

@Path(value = "/task")
@Component
public class TaskResource {

    private final TaskService taskService;

    @Autowired
    public TaskResource(TaskService taskService){
        this.taskService = taskService;
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<TaskDto> find(@QueryParam(value = "id") Long id,
                              @QueryParam(value = "title") String title,
                              @QueryParam(value = "completed") Boolean completed){
        return taskService.find(id, title, completed);
    }

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    public Response save(TaskDto taskDto){
        try {
            TaskValidator.validateNewTask(taskDto);
            return Response
                    .ok(taskService.save(taskDto.getTitle(), taskDto.getCompleted()))
                    .build();
        } catch (IllegalTaskException e) {
            return Response.status(Response.Status.NOT_FOUND).entity(e.getMessage()).build();
        }
    }

    @PUT
    @Consumes(MediaType.APPLICATION_JSON)
    public Response update(TaskDto taskDto){
        try {
            TaskValidator.validateUpdateTask(taskDto);
            taskService.update(taskDto.getId(), taskDto.getTitle(), taskDto.getCompleted());
            return Response.ok().build();
        } catch (IllegalTaskException e){
            return Response.status(Response.Status.NOT_FOUND).entity(e.getMessage()).build();
        }
    }

    @DELETE
    public Response remove(@QueryParam(value = "id") Long id){
        if(id == null){
            taskService.removeAll();
        } else {
            taskService.remove(id);
        }
        return Response.ok().build();
    }
}
