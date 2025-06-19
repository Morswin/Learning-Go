<script setup lang="ts">
    import { ref, defineProps } from 'vue';
    import { Task } from '@/libs/Task.ts';

    interface Props {
        task: Task;
    }

    const props = defineProps<Props>();

    console.log(props.task);
</script>

<template>
    <fieldset>
        <legend>
            <h3>{{ props.task.title }}</h3>
        </legend>
        <p>{{ props.task.description }}</p>
        <button>Finish</button>
        <button>Cancel</button>
    </fieldset>
</template>