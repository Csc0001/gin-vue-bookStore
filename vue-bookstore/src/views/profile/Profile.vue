<template>
  <h3>hello,{{userInfo.name}}</h3>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'userProfile',
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>

<style>
</style>
