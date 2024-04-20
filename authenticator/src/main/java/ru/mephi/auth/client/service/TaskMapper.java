package ru.mephi.auth.client.service;

import org.springframework.stereotype.Service;
import ru.mephi.auth.client.dto.TaskDto;
import ru.mephi.auth.client.entity.Task;

@Service
public class TaskMapper {
    public static TaskDto map(Task task){
        return new TaskDto(task.getId(), task.getTitle(), task.getCompletionStatus());
    }
}
