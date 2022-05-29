<template>
  <v-card class="mx-auto" max-width="600">
    <v-container fluid>
      <v-row dense>
        <v-col v-for="card in cards" :key="card.title" :cols="card.flex">
          <v-card :color="card.bgColor" height="150px">
            <v-card-title class="text-h5" :style="{ color: card.titleColor }">{{
              card.title
            }}</v-card-title>
            <v-card-subtitle
              :style="{
                color: card.subtitleColor,
                'font-size': card.subtitleSize + 'px',
              }"
              >{{ card.subtitle }}</v-card-subtitle
            >
            <v-card-text
              class="text-h2"
              :style="{
                color: card.valueColor,
              }"
              >{{ card.value }}
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import dayjs from 'dayjs'
const buddhistEra = require('dayjs/plugin/buddhistEra')
dayjs.extend(buddhistEra)

export default {
  name: 'IndexPage',
  data: () => ({
    cards: [],
  }),
  async created() {
    await this.getToday()
  },
  methods: {
    async getToday() {
      const res = await this.$axios.get('/today', {})
      const { data } = res.data
      console.log('data', data)

      // new_case: 3649
      // new_case_excludeabroad: 3647
      // new_death: 24
      // new_recovered: 5622
      // total_case: 4442648
      // total_case_excludeabroad: 4417684
      // total_death: 29972
      // total_recovered: 4367939
      // txn_date: "2022-05-29"
      // update_date: "2022-05-29 07:24:03"

      this.cards = [
        {
          title: 'Today',
          subtitle: dayjs(data.update_date).format('DD MMMM BBBB HH:mm:ss'),
          subtitleSize: 22,
          bgColor: 'white',
          flex: 12,
        },
        {
          title: 'New case',
          value: this.formatNumber(data.new_case),
          titleColor: 'white',
          valueColor: 'white',
          bgColor: '#DC5A68',
          flex: 6,
        },
        {
          title: 'Total case',
          value: this.formatNumber(data.total_case),
          titleColor: 'white',
          valueColor: 'white',
          bgColor: '#DC5A68',
          flex: 6,
        },
        {
          title: 'Death',
          value: this.formatNumber(data.new_death),
          bgColor: '#FFA500',
          titleColor: 'white',
          valueColor: 'white',
          flex: 6,
        },
        {
          title: 'Total Death',
          value: this.formatNumber(data.total_death),
          bgColor: '#FFA500',
          titleColor: 'white',
          valueColor: 'white',
          flex: 6,
        },
        {
          title: 'New recovery',

          value: this.formatNumber(data.new_recovered),
          titleColor: 'white',
          valueColor: 'white',
          bgColor: 'green',
          flex: 6,
        },
        {
          title: 'Total recovery',
          value: this.formatNumber(data.total_recovered),
          titleColor: 'white',
          valueColor: 'white',
          bgColor: 'green',
          flex: 6,
        },
      ]
    },
    formatNumber(v) {
      return v.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
    },
  },
}
</script>

<style lang="scss" scoped></style>
