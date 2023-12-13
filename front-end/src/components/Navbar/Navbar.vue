<template>
  <div>
    <div id="navbar-placeholder"></div>
    <div id="navbar">
      <nav class="bg-gray-900 text-white p-4">
        <div class="container mx-auto flex justify-between items-center">
          <div class="text-2xl font-bold">
            <RouterLink
              to="/"
              class="hover:bg-gray-600 px-4 py-2 rounded"
              active-class="bg-gray-600"
            >
              Accueil
            </RouterLink>
          </div>

          <div class="flex space-x-4">
            <a href="#Contact" class="hover:bg-gray-600 px-4 py-2 rounded">
              Contact
            </a>
            <RouterLink
              to="/adminview"
              class="hover:bg-gray-600 px-4 py-2 rounded"
              active-class="bg-gray-600"
            >
              
            </RouterLink>
          </div>
        </div>
      </nav>
    </div>
    <div id="content">
      <RouterView />
    </div>
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>
