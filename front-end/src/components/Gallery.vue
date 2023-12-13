<template>
  <div class="bg-slate-200">
    <div
      v-for="(item, index) in imageFormList"
      :key="index"
      class="pt-12 md:pt-32"
    >
      <div>
        <div class="text-center mt-8 md:mt-20">
          <h1 class="font-semibold text-2xl md:text-5xl">
            {{ item.formData.title }}
          </h1>
          <p class="font-semibold mt-2 md:mt-5">{{ item.description }}</p>
        </div>
        <div
          v-if="item.visible"
          class="flex flex-col md:flex-row"
          data-aos="fade-right"
          data-aos-delay="300"
        >
          <div class="md:w-3/5 mt-4 md:mt-10 md:ml-5">
            <!-- Utilisation des classes Tailwind pour définir la taille de la vidéo -->
            <div class="relative w-full h-0" style="padding-bottom: 56.25%">
              <video
                controls
                class="absolute inset-0 w-full h-full"
                :src="item.videoUrl"
                type="video/mp4"
              >
                Votre navigateur ne prend pas en charge la balise vidéo.
              </video>
            </div>
          </div>
          <div class="md:w-2/5 mt-1 md:pt-16 md:pl-6">
            <form
              v-if="item.visible"
              id="form"
              @submit.prevent="submitAllForms(index)"
              class="shadow-xl rounded-lg bg-slate-100 md:mr-5 pb-5 pl-5 pr-5 pt-5"
              :class="{ 'opacity-50': !item.canSubmit }"
            >
              <label for="uiuxQuestion" class="block font-semibold mt-4">{{
                item.formData.Question
              }}</label>
              <label for="comment" class="block font-semibold mt-5"
                >Commentaire :</label
              >
              <textarea
                v-model="item.formData.answer"
                id="comment"
                class="rounded-lg border p-20 w-full"
              ></textarea>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Bouton pour envoyer tous les formulaires -->
    <div class="flex justify-center mt-8">
      <button
        @click="submitAllForms"
        class="mt-4 bg-blue-500 text-white rounded-lg px-4 py-2"
      >
        Envoyer tous les formulaires
      </button>
    </div>
  </div>
</template>

<script setup>
import AOS from "aos";
import "aos/dist/aos.css";
import Swal from "sweetalert2";
import { ref } from "vue";
AOS.init();

const imageFormList = ref([
  {
    description: "Description du formulaire 1.",
    videoUrl: "/public/videos/test.mp4",
    formData: {
      title: "Formulaire 1",
      Question: "Question sur l'UI/UX test1 :",
      Answer: ""
    },
    canSubmit: true,
    visible: true
  },
  {
    description: "Description du formulaire 2.",
    videoUrl: "/public/videos/Enregistrement-LIBERTEX.mov",
    formData: {
      title: "Formulaire 2",
      Question: "Question sur l'UI/UX :",
      answer: ""
    },
    canSubmit: true,
    visible: true
  }
  // ... (autres éléments avec visible: true)
]);

const submitAllForms = async (index) => {
  try {
    // Vérifier si le formulaire a déjà été soumis en consultant le localStorage
    const isFormSubmitted = localStorage.getItem("formSubmitted");

    if (isFormSubmitted) {
      // Afficher une alerte indiquant que le formulaire a déjà été soumis
      Swal.fire({
        icon: "error",
        title: "Impossible de répondre",
        text: "Vous avez déjà répondu à ce sondage. Vous ne pouvez pas répondre une deuxième fois."
      });
      return;
    }

    const responses = imageFormList.value.map((item) => ({
      answer: item.formData.answer,
      Question: item.formData.Question
    }));

    const response = await fetch("http://localhost:3001/question", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(responses)
    });

    if (!response.ok) {
      throw new Error(
        `Erreur HTTP : ${response.status} - ${response.statusText}`
      );
    }

    const responseData = await response.json();

    if (responseData.error) {
      throw new Error(responseData.error);
    }

    // Désactiver tous les formulaires après l'envoi
    imageFormList.value.forEach((item, i) => {
      item.canSubmit = false;
      item.visible = i === index ? false : item.visible; // Masquer le formulaire actuel
    });

    // Stocker l'état de l'envoi du formulaire dans le localStorage
    localStorage.setItem("formSubmitted", "true");

    console.log("Réponse du backend :", responseData.message);

    // Afficher l'alerte pour l'envoi réussi
    Swal.fire({
      icon: "success",
      title: "Envoyé avec succès",
      text: "Vos réponses ont été soumises avec succès. Vous ne pouvez pas répondre à ce sondage une deuxième fois."
    });
  } catch (error) {
    if (error.message.includes("déjà répondu")) {
      // Afficher une alerte pour informer que l'utilisateur ne peut pas répondre une deuxième fois
      Swal.fire({
        icon: "error",
        title: "Impossible de répondre",
        text: "Vous avez déjà répondu à ce sondage. Vous ne pouvez pas répondre une deuxième fois."
      });
    } else {
      console.error("Erreur lors de l'envoi des formulaires :", error.message);
    }
  }
};
</script>
