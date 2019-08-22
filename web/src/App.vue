<template>
  <v-app>
    <AppBar/>
    <v-content>
      <BtcForm/>
    </v-content>
    <Footer/>
  </v-app>
</template>

<script>
import AppBar from './components/AppBar';
import BtcForm from './components/BtcForm';
import Footer from './components/Footer';

export default {
  name: 'App',
  components: {
    AppBar,
    BtcForm,
    Footer,
  },
  data: () => ({
  }),
};
</script>
