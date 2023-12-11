<template>
  <div>
    <div v-for="(item, index) in imageFormList" :key="index">
      <div class="text-center mt-10">
        <h1 class="font-semibold text-5xl">{{ item.title }}</h1>
        <p class="font-semibold mt-5">{{ item.description }}</p>
      </div>
      <div class="flex">
        <div class="w-1/2 mt-10 ml-5 h-screen">
          <!-- Utilisez la balise <video> pour afficher la vidéo -->
          <video controls width="100%" height="auto">
            <source :src="item.videoUrl" type="video/mp4" />
            Votre navigateur ne prend pas en charge la balise vidéo.
          </video>
        </div>
        <div class="w-1/2 pt-16 pl-5">
          <form
            @submit.prevent="submitForm(index)"
            class="shadow-xl rounded-lg bg-slate-100 mr-5 pb-5 pl-4 pr-5 pt-5"
          >
            <label for="uiuxQuestion" class="block font-semibold mt-4"> {{ item.uiuxQuestion }}</label>

            <label for="email" class="block font-semibold mt-4">Email :</label>
            <input
              v-model="item.formData.email"
              type="email"
              id="email"
              class="rounded-lg border p-2 w-full"
            />

            <label for="comment" class="block font-semibold mt-4"
              >Commentaire :</label
            >
            <textarea
              v-model="item.formData.comment"
              id="comment"
              class="rounded-lg border p-2 w-full"
            ></textarea>

            <button
              type="submit"
              class="mt-4 bg-blue-500 text-white rounded-lg px-4 py-2"
            >
              Envoyer
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const imageFormList = ref([
  {
    title: "Formulaire 1",
    description: "Description du formulaire 1.",
    videoUrl: "/public/videos/test.mp4",
    uiuxQuestion: "Question sur l'UI/UX test1 :",
    formData: {
      email: "",
      comment: ""
    }
  },
  {
    title: "Formulaire 2",
    description: "Description du formulaire 2.",
    videoUrl: "/public/videos/test.mp4",
    uiuxQuestion: "Question sur l'UI/UX :",
    formData: {
      email: "",
      comment: ""
    }
  }
  // ... ajoutez autant d'éléments que nécessaire
]);

const submitForm = async (index) => {
  try {
    const item = imageFormList.value[index];
    const response = await fetch("http://exemple.com/api/form-submit", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(item.formData)
    });

    if (!response.ok) {
      throw new Error(
        `Erreur HTTP : ${response.status} - ${response.statusText}`
      );
    }

    // Gérez la réponse du backend ici (optionnel)
    const responseData = await response.json();
    console.log("Réponse du backend :", responseData);

    // Réinitialisez seulement les champs du formulaire après l'envoi avec succès
    item.formData.uiuxQuestion = "";
    item.formData.email = "";
    item.formData.comment = "";
  } catch (error) {
    console.error("Erreur lors de l'envoi du formulaire :", error);
  }
};
</script>
