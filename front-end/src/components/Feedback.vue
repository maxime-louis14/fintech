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
              v-model="formData.nom"
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
            v-model="formData.reponse"
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
import Swal from "sweetalert2";

const formData = ref({
  nom: "",
  email: "",
  reponse: ""
});

const submitForm = async () => {
  // Vérifier si l'utilisateur a déjà envoyé un feedback
  const hasSentFeedback = localStorage.getItem("feedbackSent");

  if (hasSentFeedback && hasSentFeedback === formData.value.email) {
    // Afficher une alerte SweetAlert2 indiquant que le feedback a déjà été envoyé
    Swal.fire({
      icon: "warning",
      title: "Attention",
      text: "Vous avez déjà envoyé un feedback. Un seul feedback est autorisé."
    });
    return; // Arrêter le processus si le feedback a déjà été envoyé
  }

  try {
    const response = await fetch("http://localhost:3001/feedback", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData.value)
    });

    // Indiquer que le feedback a été envoyé dans le stockage local
    localStorage.setItem("feedbackSent", formData.value.email);

    if (response.ok) {
      // Réinitialisez le formulaire après l'envoi avec succès
      formData.value = { nom: "", email: "", reponse: "" };

      // Afficher une alerte SweetAlert2 pour le succès
      Swal.fire({
        icon: "success",
        title: "Succès",
        text: "Feedback envoyé avec succès!"
      });
    } else {
      console.error(
        `Erreur HTTP : ${response.status} - ${response.statusText}`
      );

      // Afficher une alerte SweetAlert2 pour l'échec
      Swal.fire({
        icon: "error",
        title: "Erreur",
        text: "Erreur lors de l'envoi du feedback. Veuillez réessayer."
      });
    }
  } catch (error) {
    console.error("Erreur lors de l'envoi du feedback :", error);

    let errorMessage = "Une erreur s'est produite. Veuillez réessayer.";

    // Si la réponse est une erreur 400, essayez de récupérer le message d'erreur détaillé
    if (error.response && error.response.status === 400) {
      try {
        const errorData = await error.response.json();
        if (errorData && errorData.message) {
          errorMessage = errorData.message;
        }
      } catch (jsonError) {
        console.error(
          "Erreur lors de la conversion du JSON de l'erreur :",
          jsonError
        );
      }
    }

    // Afficher une alerte SweetAlert2 pour l'échec avec le message détaillé
    Swal.fire({
      icon: "error",
      title: "Erreur",
      text: errorMessage
    });
  }
};
</script>
