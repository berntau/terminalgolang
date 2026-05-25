<script setup>
import Terminal from './components/Terminal.vue'
</script>

<template>
  <main class="app-shell">
    <div class="matrix-rain" aria-hidden="true">
      <span>01001100</span>
      <span>10110101</span>
      <span>dev</span>
      <span>01100010</span>
      <span>api</span>
      <span>11001001</span>
      <span>vue</span>
      <span>00110110</span>
      <span>go</span>
      <span>11100100</span>
      <span>json</span>
      <span>01011011</span>
    </div>

    <section class="intro-panel" aria-label="Resumo">
      <h1>curriculo bernardo dev <span>- online</span></h1>
      <p>
        Digite comandos como <code>help</code>, <code>about</code>,
        <code>skills</code>, <code>experience</code>, <code>education</code>
        ou <code>contact</code>.
      </p>
    </section>

    <Terminal />
  </main>
</template>
