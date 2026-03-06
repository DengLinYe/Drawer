<template>
  <div class="home">
    <section class="welcome">
      <div class="welcome-content">
        <img src="../assets/Drawer.png" alt="Drawer Logo" class="drawer-logo" />
        <div class="title-group">
          <h1 class="drawer-title" data-text="Drawer">Drawer</h1>
          <div class="divider"></div>
          <p class="drawer-description" data-text="Curating life's every moment.">Curating life's every moment.</p>
        </div>
      </div>
    </section>

    <section class="drawers"></section>
  </div>
</template>

<script setup>

</script>

<style>
/* 背景 */
.welcome {
  height: 100vh;

  display: flex;

  justify-content: center;
  align-items: center;

  background-color: #0f172a;
  background-image:
    radial-gradient(circle at 50% 50%, rgba(47, 58, 76, 0.8) 0%, rgba(15, 23, 42, 1) 100%),
    linear-gradient(rgb(71, 71, 79) 1px, transparent 1px),
    linear-gradient(90deg, rgb(71, 71, 79) 1px, transparent 1px);
  background-size: 100% 100%, 40px 40px, 40px 40px;

  overflow: hidden;
}

/* 欢迎内容 */
.welcome-content {
  width: 90vw;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 5vw;
}

/* Logo部分 */
.drawer-logo {
  width: 15vw;
  height: auto;

  filter: drop-shadow(20px 20px 15px rgba(0, 0, 0, 0.6)) drop-shadow(-2px -2px 10px rgba(255, 255, 255, 0.05));

  animation: logo-enter-desktop 2s cubic-bezier(0.2, 0.8, 0.2, 1) forwards;
}

@keyframes logo-enter-desktop {
  0% {
    transform: perspective(500px) translateZ(150px) translateX(15vw);
    opacity: 0;
  }

  100% {
    transform: perspective(500px) translateZ(20px) translateX(0);
    opacity: 1;
  }
}

/* 文字部分 */
.title-group {
  display: flex;
  flex-direction: column;

  opacity: 0;
  animation: text-enter-desktop 1.5s cubic-bezier(0.2, 0.8, 0.2, 1) forwards;
}

@keyframes text-enter-desktop {
  0% {
    transform: translateX(-3vw);
    opacity: 0;
  }

  40% {
    transform: translateX(-1vw);
    opacity: 0;
  }

  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

/* 标题 */
.drawer-title {
  position: relative;
  margin: 0;
  line-height: 1;
  font-size: 8vw;
  font-weight: 900;
  font-family: 'Arial Black', sans-serif;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: transparent;
  white-space: nowrap;
}

/* 标题黑底 */
.drawer-title::before {
  content: attr(data-text);
  position: absolute;
  left: 6px;
  top: 6px;
  z-index: 1;
  color: #000;
}

/* 标题文字 */
.drawer-title::after {
  content: attr(data-text);
  position: absolute;
  left: 0;
  top: 0;
  z-index: 2;
  background-image: url('../assets/background_word.jpg');
  background-size: cover;
  background-position: center;
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

/* 分隔线 */
.divider {
  width: 25vw;
  height: 4px;
  border-radius: 10px;
  background: linear-gradient(90deg, #FF7E5F, #00C9FF);
  margin: 10px 0 15px 0;
}

/* 描述 */
.drawer-description {
  position: relative;
  margin: 0;
  font-size: 2vw;
  font-weight: 900;
  font-family: 'Arial Black', sans-serif;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: transparent;
  white-space: nowrap;
}

/* 描述黑底 */
.drawer-description::before {
  content: attr(data-text);
  position: absolute;
  left: 3px;
  top: 3px;
  z-index: 1;
  color: #000;
}

/* 描述文字 */
.drawer-description::after {
  content: attr(data-text);
  position: absolute;
  left: 0;
  top: 0;
  z-index: 2;
  background-image: url('../assets/background_word.jpg');
  background-size: cover;
  background-position: center;
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

/* 适配 */
@media (max-width: 768px) {
  .welcome-content {
    flex-direction: column;
    gap: 8vw;
  }

  .drawer-logo {
    width: 30vw;

    animation: logo-enter-mobile 1.2s cubic-bezier(0.2, 0.8, 0.2, 1) forwards;
  }

  @keyframes logo-enter-mobile {
    0% {
      transform: perspective(500px) translateZ(150px) translateY(15vh);
      opacity: 0;
    }

    100% {
      transform: perspective(500px) translateZ(20px) translateY(0);
      opacity: 1;
    }
  }

  .title-group {
    align-items: center;
    text-align: center;

    animation: text-enter-mobile 1s cubic-bezier(0.2, 0.8, 0.2, 1) forwards;
    animation-delay: 0.6s;
  }

  @keyframes text-enter-mobile {
    0% {
      transform: translateY(-5vh);
      opacity: 0;
    }

    100% {
      transform: translateY(0);
      opacity: 1;
    }
  }

  .drawer-title {
    font-size: 15vw;
  }

  .divider {
    width: 60vw;
  }

  .drawer-description {
    font-size: 4vw;
  }
}
</style>