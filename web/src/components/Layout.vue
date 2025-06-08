<template>
  <div class="flex flex-row flex-nowrap">
    <div class="flex w-52 overflow-scroll">
      <SideNavbar />
    </div>
    <div class="flex w-full overflow-scroll pl-10 pr-0">
      <router-view />
    </div>
  </div>
</template>

<script setup>

import SideNavbar from "@/components/SideNavbar.vue"

</script>
