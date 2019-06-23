<template>

  <div class='container'>
    <form @submit.prevent="submitCommentForm">
      <input id="commentId" name="commentId" type="hidden" :value="comment.id"/>
      <input v-if="this.actionType === 'edit'" id="userId" name="userId" type="hidden" :value="comment.userId"/>
      <b-field
        :type="errors.has('content') ? 'is-danger': ''"
        :message="errors.has('content') ? errors.first('content') : ''"
      >
        <b-input
          v-model="comment.content"
          id="content"
          name="content"
          :placeholder="contentPlaceholder"
          type="textarea"
          min="0"
          data-vv-as="content"
          ref="content"
          v-validate="{required: true}"
        ></b-input>
      </b-field>
       <b-button tag="input"
                native-type="submit"
                :value="saveButtonName" />
    </form>
    <div class="br is-clearfix"></div>
  </div>
</template>

<script>
  import i18next from 'i18next'
  export default {
    components: {},
    props: {
      actionType: {
        type: String,
        required: true
      },
      comment: {
        required: true
      },
      newAction: {
        required: true
      },
      createAction: {
        required: true
      },
      updateAction: {
        required: true
      }
    },
    data () {
      return {
        selectExtended: true,
        saveButtonName: i18next.t('comment.form.new.submit'),
        contentPlaceholder: i18next.t('comment.form.placeholder'),
        previousContent: ''
      }
    },
    methods: {
      submitCommentForm (submitEvent) {
        this.$validator.validate().then(result => {
          if (result) {
            let articleForm = submitEvent.target
            let formData = new FormData(articleForm)
            for (var pair of formData.entries()) {
              console.log(`${pair[0]}: ${pair[1]}`)
            }
            switch (this.actionType) {
              case 'new':
                if (this.previousContent != formData.get('content')) {
                  this.previousContent = formData.get('content')
                  this.createAction(formData).then(
                    () => {
                      this.newAction()
                    }
                  )
                }
                break
              case 'edit':
                this.updateAction(formData)
                  .then(
                    () => {
                      this.newAction()
                    }
                  )
                break
            }
            return
          }
        })
      },
      updateActionType() {
        if (this.actionType === 'edit') {
          this.saveButtonName = i18next.t('comment.form.edit.submit')
          this.$refs.content.focus()
        } else {
          this.saveButtonName = i18next.t('comment.form.new.submit')
        }
      }
    },
    computed: {
    },
    updated () {
      this.updateActionType()
    },
    mounted () {
    }
  }

</script>