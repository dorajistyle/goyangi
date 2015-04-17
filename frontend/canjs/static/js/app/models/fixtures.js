/*global define*/
define(['can', 'app/models/model', 'app/models/api-path', 'settings', 'can/util/fixture'],
    function (can, Model, API, settings) {
    'use strict';

    /**
     * Item
     * @author dorajiarticle
     * @namespace article
     */

    var ARTICLES = [
  {
    id: 1,
    title: 'fashion-bookmark',
    url: 'http://fashion-bookmark.hanpage.net',
    imageName: settings.staticPath+'/images/1.jpg',
    thumbnailName: settings.staticPath+'/images/1t.jpg',
    description: 'blabla블라.',
    prev: {id: 4,
    title: 'hanpage.net',
    url: 'http://hanpage.net',
    imageName: settings.staticPath+'/images/4.jpg',
    thumbnailName: settings.staticPath+'/images/4t.jpg',
    description: 'blabla블라.',
    category: '0'
    },
    next: {id: 2,
        title: 'aynakang',
        url: 'http://aynakang.hanpage.net',
        imageName: settings.staticPath+'/images/2.jpg',
        thumbnailName: settings.staticPath+'/images/2t.jpg',
        description: 'blabla블라.',
        category: '0'},

    category: '0'
  },
  {
    id: 2,
    title: 'aynakang',
    url: 'http://aynakang.hanpage.net',
    imageName: settings.staticPath+'/images/2.jpg',
    thumbnailName: settings.staticPath+'/images/2t.jpg',
    description: 'blabla블라.',
    category: '0'
  },
  {
    id: 3,
    title: 'aile-company',
    url: 'http://aile-company.hanpage.net',
    imageName: settings.staticPath+'/images/3.jpg',
    thumbnailName: settings.staticPath+'/images/3t.jpg',
    description: 'blabla블라.',
    category: '0'
  },
  {
    id: 4,
    title: 'hanpage.net',
    url: 'http://hanpage.net',
    imageName: settings.staticPath+'/images/4.jpg',
    thumbnailName: settings.staticPath+'/images/4t.jpg',
    description: 'blabla블라.',
    category: '0'
  }

];

  var USERS = [
  {
    id: 1,
    email: "test00@test.com",
    username: "thomas00",
    name: "thomas",
    gender: 0,
    about: "ww",
    active: false,
    createdAt: "2014-11-21T07:56:55Z",
    updatedAt: "0001-01-01T00:00:00Z",
    deletedAt: "0001-01-01T00:00:00Z",
    lastLoginAt: "0001-01-01T00:00:00Z",
    currentLoginAt: "0001-01-01T00:00:00Z",
    lastLoginIp: "127.0.0.1",
    currentLoginIp: "127.0.0.1",
    Items: null
  },
  {
    id: 2,
    email: "techstilus@test.com",
    username: "techstilus",
    name: "tech",
    gender: 1,
    about: "blah",
    active: false,
    createdAt: "2014-11-21T07:56:55Z",
    updatedAt: "0001-01-01T00:00:00Z",
    deletedAt: "0001-01-01T00:00:00Z",
    lastLoginAt: "0001-01-01T00:00:00Z",
    currentLoginAt: "0001-01-01T00:00:00Z",
    lastLoginIp: "127.0.0.1",
    currentLoginIp: "127.0.0.1",
    Items: null
  }
];
    can.fixture.on = settings.canFixture;

    // articles fixture
    can.fixture('GET '+API+'/articles', function(){
      return {articles: ARTICLES, canWrite: true, category: 100};
    });

     can.fixture('GET '+API+'/articles/{id}', function(request){
      return {article: ARTICLES[request.data.id-1]};
    });

    var id= 4;
    can.fixture('POST '+API+'/articles', function(){
      return {article: {id: (id++)}}
    });

    can.fixture('PUT '+API+'/articles/{id}', function(){
      return {};
    });

    can.fixture('DELETE '+API+'/articles/{id}', function(){
      return {};
    });

    // users fixture
    can.fixture('GET '+API+'/users', function(){
      return {users: USERS};
    });

    can.fixture('GET '+API+'/users/{id}', function(request){
      return {user: USERS[request.data.id-1]};
    });

    can.fixture('POST '+API+'/users', function(){
      return {id: (id++)}
    });

    can.fixture('PUT '+API+'/users/{id}', function(){
      return {};
    });

    can.fixture('DELETE '+API+'/users/{id}', function(){
      return {};
    });

    can.fixture('POST '+API+'/authentications', function(){
      return {id: (id++)}
    });

    can.fixture('DELETE '+API+'/authentications/{id}', function(){
      return {};
    });

    can.fixture('GET '+API+'/user/current', function(request){
      return {user: USERS[0]};
    });

    can.fixture('GET '+API+'/user/email/{email}', function(request){
      return {};
    });
});
