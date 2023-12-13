<template>
  <div class="max-w-2xl mx-auto mt-8 p-4">
    <form
      @submit.prevent="addNewForm"
      class="bg-white p-6 rounded-md shadow-md"
    >
      <label class="block mb-4">
        <span class="text-gray-700">question a poser:</span>
        <input v-model="adminForm.title" class="form-input mt-1 block w-full border-4 border-indigo-500/100" />
      </label>

      <label class="block mb-4">
        <span class="text-gray-700">Fichier:</span>
        <input
          type="file"
          @change="handleFileChange"
          class="form-input mt-1 block w-full"
        />
      </label>

      <!-- ... autres champs ... -->

      <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded-md">
        Ajouter
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";

const adminForm = ref({
  title: "",
  file: null
  // ... autres champs ...
});

const handleFileChange = (event) => {
  const file = event.target.files[0];
  adminForm.value.file = file;
};

const addNewForm = () => {
  // ... logique d'ajout au tableau imageFormList ...
  // Utilisez adminForm.value.file pour accéder au fichier sélectionné
};
</script>
