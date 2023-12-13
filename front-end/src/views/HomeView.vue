<template>
  <section>
    <Home />
  </section>
  <section>
    <Gallery />
  </section>
  <section>
    <SquareGrid />
  </section>
  <section>
    <Formulaire />
  </section>
  <footer>
    <FooterVue />
  </footer>
</template>

<script setup>
import Home from "../components/Home.vue";
import Gallery from "../components/Gallery.vue";
import Formulaire from "../components/Formulaire.vue";
import FooterVue from "../components/Footer.vue";
import SquareGrid from "../components/SquareGrid.vue";
</script>
