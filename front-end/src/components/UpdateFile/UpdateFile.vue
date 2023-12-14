<template>
  <div class="max-w-2xl mx-auto mt-8 p-4">
    <div class="mt-4">
      <button
        @click="downloadJSON"
        class="bg-green-500 text-white py-2 px-4 rounded-md mr-2"
      >
        Télécharger JSON
      </button>
      <button
        @click="downloadCSV"
        class="bg-yellow-500 text-white py-2 px-4 rounded-md mr-2"
      >
        Télécharger CSV
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// Définir une référence pour stocker les données
const apiData = ref(null);

const fetchData = async (format) => {
  try {
    const response = await fetch(
      "http://localhost:3001/api/data?format=json"
    );
    const data = await response.json();
    apiData.value = data;
  } catch (error) {
    console.error(
      "Erreur lors de la récupération des données de l'API :",
      error
    );
  }
};

const downloadFile = async (fileName, mimeType, convertFunction, format) => {
  try {
    // Vérifier si les données ont été récupérées de l'API
    if (!apiData.value) {
      // Si les données ne sont pas disponibles, récupérer les données de l'API
      await fetchData(format);
    }

    // Appeler la fonction de conversion appropriée en fonction du format
    const fileData = convertFunction(apiData.value, format);

    const blob = new Blob([fileData], { type: mimeType });
    const fileUrl = URL.createObjectURL(blob);

    // Créer un élément <a> pour le téléchargement
    const a = document.createElement("a");
    a.href = fileUrl;
    a.download = fileName;

    // Ajouter l'élément <a> au DOM, déclencher le clic et le retirer
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

    // Révoquer l'URL pour libérer la mémoire
    URL.revokeObjectURL(fileUrl);
  } catch (error) {
    console.error(`Erreur lors du téléchargement de ${fileName} :`, error);
  }
};

const convertToCSV = (data) => {
  if (!data || data.length === 0) {
    console.error("Les données CSV sont undefined ou null.");
    return "";
  }

  // Extraction des clés pour les en-têtes CSV
  const headers = Object.keys(data[0]);

  // Convertir les données en une chaîne CSV
  const csvContent = data.map((row) =>
    headers.map((header) => row[header]).join(",")
  );

  // Ajouter les en-têtes CSV
  csvContent.unshift(headers.join(","));

  // Retourner la chaîne CSV complète
  return csvContent.join("\n");
};

const convertToJSON = (data) => JSON.stringify(data);

const downloadJSON = async () => {
  const format = "json";
  await downloadFile("data.json", "application/json", convertToJSON, format);
};

const downloadCSV = async () => {
  try {
    const format = "csv";
    const response = await fetch("http://localhost:3001/api/data?format=csv");
    const blob = await response.blob();

    const url = URL.createObjectURL(blob);

    const a = document.createElement("a");
    a.href = url;
    a.download = "data.csv";

    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

    URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Erreur lors du téléchargement du fichier CSV :", error);
  }
};

// Appeler fetchData lors de la création du composant
fetchData();
</script>
