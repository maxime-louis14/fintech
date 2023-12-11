<template>
  <div class="bg-slate-200 pt-52 h-screen text-center" id="Contact">
    <div
      class="p-8 rounded-lg shadow-lg bg-slate-100 mx-auto max-w-2xl mb-10"
      data-aos="fade-up"
      data-aos-delay="300"
    >
      <h1 class="text-4xl font-bold mb-6 inline-block rounded-lg">
        Contactez-nous
      </h1>
      <form @submit.prevent="submitForm" class="text-left">
        <div class="mb-4 flex">
          <div class="w-1/2 pr-2">
            <label for="name" class="block text-gray-700 font-semibold mb-2"
              >Nom :</label
            >
            <input
              v-model="formData.name"
              type="text"
              id="name"
              name="name"
              class="w-full p-2 border rounded"
            />
          </div>
          <div class="w-1/2 pl-2">
            <label for="email" class="block text-gray-700 font-semibold mb-2"
              >Email :</label
            >
            <input
              v-model="formData.email"
              type="email"
              id="email"
              name="email"
              class="w-full p-2 border rounded"
            />
          </div>
        </div>

        <div class="mb-4">
          <label for="message" class="block text-gray-700 font-semibold mb-2"
            >Message :</label
          >
          <textarea
            v-model="formData.message"
            id="message"
            name="message"
            rows="4"
            class="w-full p-20 border rounded"
          ></textarea>
        </div>

        <button
          type="submit"
          class="bg-blue-500 text-white rounded-lg px-4 py-2 mx-auto block"
        >
          Envoyer
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const formData = ref({
  name: "",
  email: "",
  message: ""
});

const submitForm = async () => {
  try {
    const response = await fetch("http://exemple.com/api/contact", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData.value)
    });

    if (response.ok) {
      // Réinitialisez le formulaire après l'envoi avec succès
      formData.value = { name: "", email: "", message: "" };
    } else {
      console.error(
        `Erreur HTTP : ${response.status} - ${response.statusText}`
      );
    }
  } catch (error) {
    console.error("Erreur lors de l'envoi du formulaire :", error);
  }
};
</script>

<style scoped>
/* Ajoutez ici des styles spécifiques au composant si nécessaire */
</style>
