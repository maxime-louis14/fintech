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

const downloadFile = async (fileName, mimeType, convertFunction) => {
  try {
    const response = await fetch("http://localhost:3001/api/data");
    const data = await response.json();
    const fileData = convertFunction(data);
    const blob = new Blob([fileData], { type: mimeType });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = fileName;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  } catch (error) {
    console.error(`Erreur lors du téléchargement de ${fileName} :`, error);
  }
};

const convertToCSV = (data) => {
  const header = Object.keys(data[0]).join(",");
  const rows = data.map((row) => Object.values(row).join(","));
  return `${header}\n${rows.join("\n")}`;
};

const convertToJSON = (data) => JSON.stringify(data);

const convertToSQLite = (data) => {
  const db = new SQL.Database();
  const tableName = "questionnaire";

  // Create a table in the SQLite database
  db.run(
    `CREATE TABLE ${tableName} (title TEXT, email TEXT, reponse TEXT, question TEXT, createdAt TEXT)`
  );

  // Insert data into the SQLite database
  data.forEach((entry) => {
    db.run(`INSERT INTO ${tableName} VALUES (?, ?, ?, ?, ?)`, [
      entry.title,
      entry.email,
      entry.reponse,
      entry.question,
      entry.createdAt
    ]);
  });

  // Export the SQLite database to a Uint8Array
  return db.export();
};

const downloadJSON = () => {
  downloadFile("data.json", "application/json", convertToJSON);
};

const downloadCSV = () => {
  downloadFile("data.csv", "text/csv", convertToCSV);
};

const downloadSQLite = () => {
  downloadFile("data.sqlite", "application/x-sqlite3", convertToSQLite);
};
</script>
