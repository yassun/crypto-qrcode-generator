<template>
  <v-app>
    <AppBar/>
    <v-content>
      <router-view/>
    </v-content>
    <Footer/>
  </v-app>
</template>

<script>
import AppBar from './components/AppBar';
import Footer from './components/Footer';

export default {
  name: 'App',
  components: {
    AppBar,
    Footer,
  },
  data: () => ({
  }),
};
</script>
