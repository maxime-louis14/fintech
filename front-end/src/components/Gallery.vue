<template>
  <div class="bg-slate-200">
    <div
      v-for="(item, index) in imageFormList"
      :key="index"
      class="pt-12 md:pt-32"
    >
      <div class="text-center mt-8 md:mt-20">
        <h1 class="font-semibold text-2xl md:text-5xl">{{ item.title }}</h1>
        <p class="font-semibold mt-2 md:mt-5">{{ item.description }}</p>
      </div>
      <div
        class="flex flex-col md:flex-row"
        data-aos="fade-up"
        data-aos-delay="300"
      >
        <div class="md:w-3/5 mt-4 md:mt-10 md:ml-5">
          <!-- Utilisez la balise <video> pour afficher la vidéo -->
          <video controls width="100%" height="auto">
            <source :src="item.videoUrl" type="video/mp4" />
            Votre navigateur ne prend pas en charge la balise vidéo.
          </video>
        </div>
        <div class="md:w-2/5 mt-1 md:pt-16 md:pl-6">
          <form
            @submit.prevent="submitForm(index)"
            class="shadow-xl rounded-lg bg-slate-100 md:mr-5 pb-5 pl-5 pr-5 pt-5"
          >
            <label for="uiuxQuestion" class="block font-semibold mt-4">
              {{ item.uiuxQuestion }}</label
            >

            <label for="comment" class="block font-semibold mt-5"
              >Commentaire :</label
            >
            <textarea
              v-model="item.formData.answer"
              id="comment"
              class="rounded-lg border p-20 w-full"
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
import AOS from "aos";
import "aos/dist/aos.css"; // You can also use <link> for styles
AOS.init();
import { ref } from "vue";

const imageFormList = ref([
  {
    title: "Formulaire 1",
    description: "Description du formulaire 1.",
    videoUrl: "/public/videos/test.mp4",
    uiuxQuestion: "Question sur l'UI/UX test1 :",
    formData: {
      Answer: ""
    }
  },
  {
    title: "Formulaire 2",
    description: "Description du formulaire 2.",
    videoUrl: "/public/videos/test.mp4",
    uiuxQuestion: "Question sur l'UI/UX :",
    formData: {
      answer: ""
    }
  }
  // ... ajoutez autant d'éléments que nécessaire
]);

const submitForm = async (index) => {
  try {
    const item = imageFormList.value[index];
    const response = await fetch("http://localhost:3001/question", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify([{ answer: item.formData.answer }])
    });

    if (!response.ok) {
      throw new Error(
        `Erreur HTTP : ${response.status} - ${response.statusText}`
      );
    }

    // Gérez la réponse du backend ici (optionnel)
    const responseData = await response.json();

    // Vérifiez si la réponse contient des erreurs
    if (responseData.error) {
      throw new Error(responseData.error);
    }

    // Réinitialisez seulement les champs du formulaire après l'envoi avec succès
    item.formData.answer = "";

    // Affichez un message de succès ou faites d'autres actions nécessaires
    console.log("Réponse du backend :", responseData.message);
  } catch (error) {
    console.error("Erreur lors de l'envoi du formulaire :", error.message);
  }
};
</script>
