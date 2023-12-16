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
    <Feedback />
  </section>
  <footer>
    <FooterVue />
  </footer>
</template>

<script setup>
import Home from "../components/Home.vue";
import Gallery from "../components/Gallery.vue";
import Feedback from "../components/Feedback.vue";
import FooterVue from "../components/Footer.vue";
</script>
