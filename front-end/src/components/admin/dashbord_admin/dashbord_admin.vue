<script setup>
import { ref } from "vue";
import statistiques from "@/components/admin/UpdateFile/UpdateFile.vue";
import test from "@/components/test/test.vue";

const step = ref("statistiques");
</script>

<template>
  <div class="flex">
    <div class="w-64 bg-gray-50 shadow-xl p-8 rounded-lg h-screen">
      <!-- Barre latérale avec des éléments de tableau de bord -->
      <h2 class="text-xl font-bold mb-4 text-center">Tableau de bord</h2>
      <ul class="">
        <li
          class="text-black no-underline text-center font-semibold text-lg cursor-pointer rounded-full mb-5"
          @click="step = 'statistiques'"
        >
          Statistiques
        </li>
        <li
          class="text-black no-underline text-center font-semibold text-lg cursor-pointer rounded-full mb-5"
          @click="step = 'projects'"
        >
          projects
        </li>
        <li
          class="text-black no-underline text-center font-semibold text-lg cursor-pointer rounded-full mb-5"
          @click="step = 'sevices'"
        >
          Sevices
        </li>
        <li
          class="text-black no-underline text-center font-semibold text-lg cursor-pointer rounded-full mb-5"
          @click="step = 'chat'"
        >
          Chat
        </li>
        <li
          class="text-black no-underline text-center font-semibold text-lg cursor-pointer rounded-full mb-5"
          @click="step = 'paramètres'"
        >
          Paramètres
        </li>
      </ul>
    </div>
    <div>
      <statistiques v-if="step === 'statistiques'" />
    </div>
    <div>
      <test v-if="step === 'projects'" />
    </div>
    <div>
      <test v-if="step === 'sevices'" />
    </div>
    <div>
      <test v-if="step === 'chat'" />
    </div>
    <div>
      <test v-if="step === 'paramètres'" />
    </div>

    <div class="flex-1 p-8">
      <!-- Contenu principal -->
      <router-view></router-view>
    </div>
  </div>
</template>
