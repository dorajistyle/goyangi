<template>
  <div>
    <commentForm
    v-bind:actionType="formActionType"
    v-bind:comment="comment"
    :createAction="createAction"
    :updateAction="updateAction"
    :newAction="newCommentForm"
    />
    <commentList
    :commentList="commentList"
    :parentId="parentId"
    :deleteAction="deleteAction"
    :editAction="editCommentForm"
    :loadMoreAction="loadMoreAction"
    />
    <router-view></router-view>
  </div>

</template>

<script>
  import commentForm from './Form'
  import commentList from './List'

  export default {
    components: {
      commentForm,
      commentList
    },
    data () {
    return {
      formActionType: 'new',
      comment: {}
    }
  },

    props: {
      parentId: {
        requred: true
      },
      createAction: {
        required: true
      },
      updateAction: {
        required: true
      },
      deleteAction: {
        required: true
      },
      loadMoreAction: {
        required: true
      },
      commentList: {
        required: true
      }
    },
    methods: {
      newCommentForm() {
        this.formActionType = 'new'
        this.comment = {}
      },
      editCommentForm(comment) {
        this.formActionType = 'edit'
        this.comment = Object.assign({}, comment)
      }
  },
  }
</script>