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
      <button
        @click="downloadSQLite"
        class="bg-blue-500 text-white py-2 px-4 rounded-md"
      >
        Télécharger SQLite
      </button>
    </div>
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

const downloadJSON = async () => {
  try {
    const response = await fetch("http://localhost:3001/api/data");
    const data = await response.json();

    // Logique pour télécharger le fichier JSON
    const jsonData = JSON.stringify(data);
    downloadFile(jsonData, "data.json", "application/json");
  } catch (error) {
    console.error("Erreur lors du téléchargement JSON :", error);
  }
};

const downloadCSV = async () => {
  try {
    const response = await fetch("http://localhost:3001/api/data");
    const data = await response.json();

    // Logique pour télécharger le fichier CSV
    const csvData = convertToCSV(data);
    downloadFile(csvData, "data.csv", "text/csv");
  } catch (error) {
    console.error("Erreur lors du téléchargement CSV :", error);
  }
};

const downloadSQLite = async () => {
  try {
    const response = await fetch("http://localhost:3001/api/sqlite");
    const buffer = await response.arrayBuffer();

    // Logique pour télécharger le fichier SQLite
    const db = new SQL.Database(new Uint8Array(buffer));
    const sqliteData = db.export();
    downloadFile(sqliteData, "data.sqlite", "application/x-sqlite3");
  } catch (error) {
    console.error("Erreur lors du téléchargement SQLite :", error);
  }
};

const downloadFile = (data, fileName, mimeType) => {
  const blob = new Blob([data], { type: mimeType });
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = fileName;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};

const convertToCSV = (data) => {
  const header = Object.keys(data[0]).join(",");
  const rows = data.map((row) => Object.values(row).join(","));
  return `${header}\n${rows.join("\n")}`;
};
</script>
