<script setup lang="ts">
    import { ref } from 'vue';
    import { Task } from '@/libs/Task.ts'

    const tasks = ref<Task[]>([]);

    tasks.value.push(new Task("Title", "Description"));
    console.log(tasks.value);
</script>

<template>
    <article>
        <!-- Here goes the main stuff -->
        <h1>TODO App</h1>
        <TaskComponent v-for="task in tasks" :task="task"></TaskComponent>
    </article>
    <nav>
        <!-- Here go the nav inputs -->
    </nav>
</template>