<template>
  <div class="questionnaire">
    <Question
      v-for="(question, index) in questions"
      :key="index"
      :question="question"
      v-model="responses[index]"
    />
    <button @click="submitQuestionnaire">Envoyer le questionnaire</button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Question from './Question.vue';
import questionsData from '@/data/questions.json';

const questions = ref(questionsData.questions);
const responses = ref([]);

const submitQuestionnaire = () => {
  console.log('Réponses :', responses.value);
};
</script>
