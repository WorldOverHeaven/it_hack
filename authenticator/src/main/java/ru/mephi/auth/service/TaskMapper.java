package ru.mephi.auth.service;

import org.springframework.stereotype.Service;
import ru.mephi.auth.dto.TaskDto;
import ru.mephi.auth.entity.Task;

@Service
public class TaskMapper {
    public static TaskDto map(Task task){
        return new TaskDto(task.getId(), task.getTitle(), task.getCompletionStatus());
    }
}
