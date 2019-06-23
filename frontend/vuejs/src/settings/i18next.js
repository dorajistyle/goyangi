const en = require('@/locales/en/translation.json')
const ko = require('@/locales/ko/translation.json')

export default {debug: true,
  fallbackLng: 'en',
  resources: {
    en: { translation: en },
    ko: { translation: ko }
  }
}
