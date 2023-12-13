<template>
  <div class="question p-4 border-b mb-4">
    <label class="font-semibold mb-2">{{ question.text }}</label>

    <template v-if="question.type === 'text'">
      <input
        v-model="response"
        class="p-2 border rounded focus:outline-none focus:ring focus:border-blue-300"
        @input="emitResponse"
        placeholder="Saisissez votre réponse"
      />
    </template>

    <template v-else-if="question.type === 'checkbox'">
      <div
        v-for="(option, index) in question.options"
        :key="index"
        class="mb-2"
      >
        <input
          type="checkbox"
          v-model="response"
          :value="option"
          class="mr-2 focus:outline-none focus:ring focus:border-blue-300"
          @change="emitResponse"
        />
        {{ option }}
      </div>
    </template>

    <template v-else-if="question.type === 'radio'">
      <div
        v-for="(option, index) in question.options"
        :key="index"
        class="mb-2"
      >
        <input
          type="radio"
          v-model="response"
          :value="option"
          class="mr-2 focus:outline-none focus:ring focus:border-blue-300"
          @change="emitResponse"
        />
        {{ option }}
      </div>
    </template>

    <template v-else-if="question.type === 'dropdown'">
      <select
        v-model="response"
        class="p-2 border rounded focus:outline-none focus:ring focus:border-blue-300"
        @change="emitResponse"
      >
        <option value="" disabled selected>Sélectionnez une option</option>
        <option
          v-for="(option, index) in question.options"
          :key="index"
          :value="option"
        >
          {{ option }}
        </option>
      </select>
    </template>

    <template v-else-if="question.type === 'textarea'">
      <textarea
        v-model="response"
        class="p-2 border rounded focus:outline-none focus:ring focus:border-blue-300"
        @input="emitResponse"
        placeholder="Saisissez votre réponse"
      ></textarea>
    </template>

    <!-- Ajoutez d'autres types de champs selon les besoins -->
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits } from "vue";

const { question, modelValue } = defineProps(["question", "modelValue"]);
const response = ref(modelValue);

// Émettre la réponse vers le composant parent
const emitResponse = () => {
  if (response.value !== null) {
    emit("updateResponse", response.value);
  }
};
</script>
