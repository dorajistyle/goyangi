define(['can', 'app/models/user/user', 'app/models/user/user-current', 'app/models/user/follower', 'app/models/user/following',
        'app/partial/liking', 'app/models/user/liking', 'app/models/user/liked',
        'util', 'i18n', 'jquery'],
    function (can, User, UserCurrent, Follower, Following, Liking, LikingModel, LikedModel, util, i18n, $) {
    'use strict';


    /**
     * Instance of User Contorllers.
     * @private
     */
    var read;

    /**
     * Control for User Profile
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#Read
     * @constructor
     */
    var Read = can.Control.extend({
        init: function () {
            util.logInfo('*User/Read', 'Initialized');
        },


        load: function (id) {
            User.findOne({id: id}, function (userData) {
              var templateData = new can.Observe({
                user: userData.user,
                isAuthor: userData.isAuthor,
                currentUserId: userData.currentUserId,
                parentId: userData.user.id,
                likingList: userData.user.likingList
              });
                read.userData = templateData;
                // Follower.findAll({id: id}, function (followers) {
                    // read.followersData = followers;
                    // Following.findAll({id: id}, function (following) {
                        // read.followingData = following;
                        can.when(read.show()).then(function () {
                          util.refreshTitle();
                          new Liking($(read.element), {parentId: read.userData.user.id, parentName: userData.user.username,
                            view: read, model: LikingModel, modelLiked: LikedModel, likePrefix: "profile.follow", unlikePrefix: "profile.unfollow"});
                        });
                    // })
                // });
            });
        },
        /**
         * Show profile view.
         * @memberof users#Read
         */
        show: function () {
            util.logJson('*User/Read.show data', read.userData);
            util.logDebug('*User/Read.show data', read.userData.user.email);


            // this.following = this.followingData.following;
            // this.followers = this.followersData.followers;
            // this.followingFlag = {hasNext: this.followingData.hasNext, currentPage: this.followingData.currentPage};
            // this.followersFlag = {hasNext: this.followersData.hasNext, currentPage: this.followersData.currentPage};
            this.element.html(can.view('views_user_profile_stache', {
                userData: read.userData
                // followingFlag: this.followingFlag,
                // following: this.following,
                // followersFlag: this.followersFlag,
                // followers: this.followers
            }));
        },
        '.perform-follow click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performFollow();
            return false;
        },

        '.perform-unfollow click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performUnfollow();
            return false;
        },
        '.show-more-following click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.showMoreFollowing();
            return false;
        },
        '.show-more-followers click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.showMoreFollowers();
            return false;
        },
        performFollow: function () {
            var form = this.element.find('#followersForm');
            var values = can.deparam(form.serialize());
            var $form = $(form);
            if (!$form.data('submitted')) {
                $form.data('submitted', true);
                var followBtn = this.element.find('.perform-unfollow');
                followBtn.attr('disabled', 'disabled');
                can.when(Follower.create(values)).then(function (result) {
                    followBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    util.showSuccessMsg(i18n.t('profile.follow.done', result.email));
                    read.load(values.id);
                }, function (xhr) {
                    util.handleStatusWithErrorMsg(xhr, i18n.t('profile.follow.fail'));
                });
            }
        },
        performUnfollow: function () {
            var form = this.element.find('#followersForm');
            var values = can.deparam(form.serialize());
            var $form = $(form);
            if (!$form.data('submitted')) {
                $form.data('submitted', true);
                var followBtn = this.element.find('.perform-follow');
                followBtn.attr('disabled', 'disabled');
                can.when(Follower.destroy(values.id)).then(function (result) {
                    followBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    util.showSuccessMsg(i18n.t('profile.unfollow.done', result.email));
                    read.load(values.id);
                }, function (xhr) {
                    util.handleStatusWithErrorMsg(xhr, i18n.t('profile.unfollow.fail'));
                });
            }
        },
        showMoreFollowing: function () {
            var moreFollowingBtn = this.element.find('.show-more-following');
            moreFollowingBtn.attr('disabled', 'disabled');
            Following.findAll({id: read.user.user.id, currentPage: read.followingData.currentPage + 1}, function (following) {
                    following.hasNext ? moreFollowingBtn.removeAttr('disabled') : moreFollowingBtn.addClass('uk-hidden');
                    read.followingData.attr('currentPage', following.currentPage);
                    read.followingData.attr('hasNext', following.hasNext);
                    read.following.attr(read.following.attr().concat(following.following.attr()));
                }
            );
        },
        showMoreFollowers: function () {
            var moreFollowersBtn = this.element.find('.show-more-followers');
            moreFollowersBtn.attr('disabled', 'disabled');
            Follower.findAll({id: read.user.user.id, currentPage: read.followersData.currentPage + 1}, function (followers) {
                    followers.hasNext ? moreFollowersBtn.removeAttr('disabled') : moreFollowersBtn.addClass('uk-hidden');
                    read.followersData.attr('currentPage', followers.currentPage);
                    read.followersData.attr('hasNext', followers.hasNext);
                    read.followers.attr(read.followers.attr().concat(followers.followers.attr()));
                }
            );
        }

    });

    /**
     * Router for profile.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name users#ProfileRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*Users/Profile/Router', 'Initialized')
        },
        'profile route': function () {
            UserCurrent.findOne({}, function (result) {
                can.route.attr({'id': result.user.id});
            });
        },
        'profile/:id route': function (data) {
            read = new Read(util.getFreshApp());
            read.load(data.id);
        }
    });

    return Router;
});
