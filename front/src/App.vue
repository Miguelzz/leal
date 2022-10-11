<template>
  <div>
    <nav class="navbar navbar-expand-lg bg-light">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">Store leal</a>

        <ul class="nav justify-content-center">
          <li class="nav-item">
            <a class="nav-link disabled" href="#">Points {{ points }}</a>
          </li>

          <li class="nav-item">
            <a class="nav-link disabled" href="#">balance {{ balance }}</a>
          </li>
        </ul>
      </div>
    </nav>

    <section class="container products mt-4">
      <div v-for="p in products" :key="p" class="card" style="width: 18rem">
        <div class="point">{{ p.price / 1000 + 5 }}</div>
        <img :src="p.image" class="card-img-top" />
        <div class="card-body">
          <h5 class="card-title">{{ p.name }} ${{ p.price }}</h5>
          <p class="card-text">
            For every purchase you make in the month of October you will have
            double points
          </p>
          <div
            class="d-flex flex-row justify-content-around align-items-center"
          >
            <a
              @click="buy(p)"
              v-if="balance >= p.price"
              class="btn btn-primary btn-success"
              >BUY</a
            >
            <a
              @click="redeem(p)"
              v-if="points >= p.price / 1000 + 5"
              class="btn btn-primary btn-warning"
              >REDEEM</a
            >
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "App",
  components: {},
  data() {
    return {
      points: 0,
      balance: 200_000,
      products: [],
    };
  },

  methods: {
    buy(p) {
      fetch("http://localhost:8080/buy/0/" + p.id, { method: "POST" })
        .then((response) => response.json())
        .then((data) => {
          this.balance -= p.price;
          this.points = data.points;
        });
    },
    redeem(p) {
      fetch("http://localhost:8080/redeem/0/" + p.id, { method: "POST" })
        .then((response) => response.json())
        .then((data) => {
          this.points = data.points;
        });
    },
  },

  mounted() {
    fetch("http://localhost:8080/products")
      .then((response) => response.json())
      .then((data) => {
        this.products = data;
      });

    fetch("http://localhost:8080/points/0")
      .then((response) => response.json())
      .then((data) => {
        this.points = data.points;
      });
  },
};
</script>

<style>
.products {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 2rem;
}

.point {
  position: absolute;
  background-color: orange;
  padding: 0.5rem;
  right: 0;
  border-radius: 7px;
}
</style>
