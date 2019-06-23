<template>
  <div :class="this.componentClass">
    <a
      v-if="this.likingList.isLiked"
      @click="this.deleteAction"
      class="button">
        <span class="icon">
          <i class="fas fa-heart"></i>
        </span>
        <span>
          {{this.likingList.count }}
        </span>
    </a>
    <a
      v-else
      @click="this.createAction"
      class="button">
        <span class="icon">
          <i class="far fa-heart"></i>
        </span>
        <span>
          {{this.likingList.count }}
        </span>
    </a>
  </div>
</template>

<script>
  export default {
    props: {
      componentClass: {
        required: false,
        default: ''
      },
      likingList: {
        required: true,
        default () {
          return {
            hasPrev: false,
            hasNext: false,
            count: 0,
            currentPage: 1,
            isLiked: false,
            likings: []
          }
        }
      },
      createAction: {
        required: true
      },
      deleteAction: {
        required: true
      }
    },
    data () {
      return {
      }
    },
    methods: {
    }
  }

</script>