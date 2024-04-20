package ru.mephi.auth.resource;

import io.micrometer.common.util.StringUtils;
import ru.mephi.auth.dto.TaskDto;
import ru.mephi.auth.exception.IllegalTaskException;

public class TaskValidator {
    public static void validateNewTask(TaskDto taskDto) {
        String title = taskDto.getTitle();
        Boolean completed = taskDto.getCompleted();

        if (title == null || StringUtils.isBlank(title)) {
            throw new IllegalTaskException("Title is not present");
        }
        if (completed == null) {
            throw new IllegalTaskException("Completed is not present");
        }
    }

    public static void validateUpdateTask(TaskDto taskDto) {
        Long id = taskDto.getId();

        if (id == null) {
            throw new IllegalTaskException("Id is not present");
        }
    }
}
