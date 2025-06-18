<script setup lang="ts">
    import { ref, defineProps } from 'vue';

    interface Props {
        title: string,
        description: string
    }

    const props = defineProps<Props>();
</script>

<template>
    <fieldset>
        <legend>
            <h3>{{ props.title }}</h3>
        </legend>
        <p>{{ props.description }}</p>
        <button>Finish</button>
        <button>Cancel</button>
    </fieldset>
</template>