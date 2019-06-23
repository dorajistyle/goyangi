<template>

  <div class='container'>
    <nav class="breadcrumb" aria-label="breadcrumbs">
      <ul>
        <li>
          <router-link to="/articles" class="has-text-centered">
            <span>
              {{ $t('article.view.item.list') }}
            </span>
          </router-link>
        </li>
        <li v-if="article.id">
          <router-link :to="`/articles/${article.id}`" class="has-text-centered">
            <span>
              {{ article.title }}
            </span>
          </router-link>
        </li>
        <li class="is-active"><a href="#" aria-current="page">{{ actionName }}</a></li>
      </ul>
    </nav>

      <form @submit.prevent="submitForm">
        <input id="id" name="id" type="hidden" :value="article.id"/>
        <b-field horizontal
          v-if="true"
          label="Title"
          :type="errors.has('title') ? 'is-danger': ''"
          :message="errors.has('title') ? errors.first('title') : ''"
        >
          <b-input
            v-model="article.title"
            id="title"
            name="title"
            data-vv-as="title"
            v-validate="{required: true}"
          >
          </b-input>
        </b-field>
        <b-field horizontal
          label="Content"
          :type="errors.has('content') ? 'is-danger': ''"
          :message="errors.has('content') ? errors.first('content') : ''"
        >
          <b-input
            v-model="article.content"
            id="content"
            name="content"
            placeholder="Number"
            type="textarea"
            min="0"
            data-vv-as="content"
            v-validate="{required: true}"
          ></b-input>
        </b-field>
        <input
          class="button is-link"
          type="submit"
          :value="saveButtonName"
        />
      </form>
  </div>
</template>

<script>
  import i18next from 'i18next'
  import { mapActions, mapGetters } from 'vuex'
  export default {
    components: {},
    computed: mapGetters({
      article: 'article'
    }),
    props: {
      action: {
        type: String,
        required: true
      }
    },
    data () {
      return {
        selectExtended: true,
        actionName: i18next.t('article.view.form.new.title'),
        saveButtonName: i18next.t('article.view.form.new.submit')
      }
    },
    methods: {
      ...mapActions([
        'newArticle',
        'createArticle',
        'updateArticle'
      ]),
      submitForm (submitEvent) {
        this.$validator.validate().then(result => {
          if (result) {
            let articleForm = submitEvent.target
            let formData = new FormData(articleForm)
            for (var pair of formData.entries()) {
              console.log(`${pair[0]}: ${pair[1]}`)
            }
            switch (this.action) {
              case 'new':
                this.createArticle(formData)
                break
              case 'edit':
                this.updateArticle(formData)
                break
            }
            return
          }
        })
      },
      showList () {
        this.$router.push({path: '/articles'})
      }
    },
    mounted () {
      if (this.action === 'new') { this.newArticle() }
      if (this.action === 'edit') {
        this.actionName = i18next.t('article.view.form.edit.title')
        this.saveButtonName = i18next.t('article.view.form.edit.submit')
        if(this.article.id == null) this.showList()
      }
    }
  }

</script>