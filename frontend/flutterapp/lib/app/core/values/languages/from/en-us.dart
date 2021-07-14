final Map<String, String> enUs = {
  "service": {"name": "Goyangi"},
  "utils": {
    "stringify": {
      "notSupport":
          "JSON stringify function supports modern browser(greater than IE7) only."
    }
  },
  "error": {
    "incorrectAccessDetected": "Incorrect access is detected.",
    "pleaseCheckTheURL": "Please Check the URL",
    "or": "or",
    "loginPlease": "Login please.",
    "adminRequired": "This function can only accessed by Admin.",
    "forbidden": "Forbidden.",
    "unauthorized": "Unauthorized.",
    "notFound": "Not found.",
    "internalServerError": "Error occured.",
    "unknown": "Unknown error. Please contact us to figure it out."
  },
  "navbar": {
    "home": "Home",
    "login": "Login",
    "logout": "Logout",
    "registration": "Registration",
    "myAccount": "My Account",
    "admin": "Admin",
    "profile": "Profile",
    "setting": "Setting",
    "aboutUs": "About Us",
    "sitemap": "Sitemap",
    "howItWorks": "How it Works",
    "languages": {"title": "Languages", "english": "English", "korean": "한글"},
    "examples": {"articles": "Articles", "upload": "Upload"}
  },
  "destroy": {
    "confirm": {
      "title": "Delete %s.",
      "content": "Do you really want to delete this %s?",
      "close": "Close",
      "submit": "Delete the %s."
    },
    "done": "%s is deleted successfully.",
    "fail": "The %s is not deleted."
  },
  "typeahead": {
    "emptyMessage": "Unable to find any item that match the current query."
  },
  "page": {"previous": "Previous", "next": "Next"},
  "main": {
    "title": "Goyangi",
    "new": "new",
    "description":
        "Goyangi is a web foundation who want to make web services with the Go server. It is also a good example of how to develop a large web application that combined fancy Go packages. Goyangi means cat in Korean."
  },
  "validation": {
    "email": "Please enter an email address.",
    "emailOr": "Please enter your email address or ",
    "content": "Please enter a content.",
    "agreement": "Please confirm the terms and conditions.",
    "password":
        "Please enter correct password. Length should be between 6 to 20 charcters.",
    "verifyPassword":
        "The verified password of the user, which must be identical to password.",
    "username":
        "Please enter a valid username. Username must be between 4 to 16 alphanumeric characters in length (a-z and/or 0-9).",
    "badUsername": "Username contains bad words.",
    "newPasswordConfirm":
        "‘New Password Confirm’ should be exactly same as ‘New Password’.",
    "contentTooShort":
        "It's too short. Content should be at least 3 characters.",
    "unauthorized": "Please log in first.",
    "unknown": "Whoops It didn't work because an unknown error."
  },
  "welcome": {
    "login": {
      "title": "Welcome to Service.",
      "content":
          "Enjoy today, this moment, right now. Present moment is just present."
    }
  },
  "registration": {
    "title": "New to Service? Register!",
    "enterUsername": "Please enter username.",
    "enterEmail": "Please enter your email address.",
    "enterPassword": "Please enter your password.",
    "enterVerifyPassword": "Please verify a password.",
    "welcome": "Welcome!",
    "already": "%s is already taken.",
    "available": "You can use %s.",
    "submit": "Register",
    "error": {"fail": "Registration failed. Please try it again."}
  },
  "login": {
    "info": "Please enter your email and password.",
    "enterEmail": "Please enter your email address.",
    "enterPassword": "Please enter your password.",
    "welcome": "Welcome!!!",
    "already": "You already logged in.",
    "userDisabled": "This account not available.",
    "rememberMe": "Remember me",
    "forgotPassword": "Forgot password",
    "submit": "Login",
    "done": "You logged in succesfully.",
    "error": {
      "userNotFound": "You didn't registered yet.",
      "passwordIncorrect": "Password is incorrect.",
      "fail": "Login failed. Please try it again."
    }
  },
  "logout": {
    "done": "Goodbye. See you again.",
    "alreadyDone": "You already logged out.",
    "error": "Logout uncompleted. Please try it again."
  },
  "passwordReset": {
    "send": {
      "email": "Please enter your email",
      "submit": "Send me the password reset token",
      "sent": {
        "done": "We sent the password reset token.",
        "fail": "Email not sent. Try it again later."
      },
      "validation": {
        "emailNotExist":
            "You are not joined yet. Check your email again or Please join us."
      }
    },
    "reset": {
      "submit": "Change my password",
      "newPassword": "New Password",
      "newPasswordConfirm": "New Password Confirm",
      "enterNewPassword": "Please enter new password.",
      "enterNewPasswordConfirm": "Please enter your new password again.",
      "updated": {
        "done": "User password is updated.",
        "fail": "User password is not updated."
      }
    },
    "error": {"tokenExpired": "Password reset token expired."}
  },
  "emailVerification": {
    "title": "Email Verification",
    "done": "Email verified.",
    "fail": "Email not verified.",
    "send": {
      "notYet": "Your Email not verified yet.",
      "submit": "Send me the email verification token again.",
      "sent": {
        "done": "We sent the email verification token.",
        "fail": "Email not sent. Try it again later."
      }
    },
    "error": {"tokenExpired": "Email verification token expired."}
  },
  "admin": {
    "title": "Admin Tools",
    "role": {
      "title": "Roles",
      "name": "Name",
      "description": "Description",
      "enterName": "Enter name",
      "enterDescription": "Enter description",
      "formTitle": "Create role",
      "submit": "Create role",
      "showNewForm": "I want to create a new Role",
      "nameValidation": "Name should not be blank.",
      "create": {
        "done": "Role is created successfully.",
        "fail": "The role is not created."
      },
      "update": {
        "title": "Update role",
        "done": "Role updated successfully.",
        "fail": "The role is not updated."
      },
      "confirm": {
        "title": "Delete Role",
        "content": "Do you really want to delete this role?",
        "close": "Close",
        "submit": "Delete the role"
      }
    },
    "user": {
      "title": "User",
      "email": "Email",
      "roles": "Roles",
      "hasRole": "--  User has __count__ role.  --",
      "hasRole_plural": "--  User has  __count__  roles.  --",
      "active": "Active",
      "formTitle": "Add Role to User",
      "searchRole": "Search Role",
      "submit": "Add the role",
      "confirmDelete": {
        "title": "Delete user.",
        "content": "Do you really want to delete this user?",
        "close": "Close",
        "submit": "Delete the user."
      },
      "role": {
        "confirmDelete": {
          "title": "Delete role from user.",
          "content": "Do you really want to delete this role from user?",
          "close": "Close",
          "submit": "Delete the role from user."
        },
        "add": {
          "done": "Role added to user successfully.",
          "fail": "The role not added to user."
        },
        "delete": {
          "done": "Role deleted from user successfully.",
          "fail": "The role is not deleted from user."
        }
      },
      "toggleActivate": {
        "done": "User's activate status changed successfully.",
        "fail": "User's active status not changed."
      },
      "delete": {
        "done": "Role deleted from user successfully.",
        "fail": "The role is not deleted from user."
      }
    }
  },
  "user": {
    "profile": "Profile",
    "setting": "Setting",
    "email": "Email",
    "id": "id",
    "view": {
      "oauth": {
        "ifYouRegister": "If you register, ",
        "provider": {
          "google": "Google",
          "github": "Github",
          "yahoo": "Yahoo",
          "facebook": "Facebook",
          "twitter": "Twitter",
          "linkedin": "Linked In",
          "kakao": "Kakao",
          "naver": "Naver"
        },
        "connect": {
          "title": "Oauth Connection",
          "close": "Close",
          "prefix": "Connect with ",
          "postfix": ""
        },
        "revoke": {
          "title": "Revoke Oauth Connection",
          "submit": "Revoke",
          "close": "Close",
          "content": {"prefix": "Are you sure to revoke from ", "postfix": "?"},
          "prefix": "Revoke from ",
          "postfix": ""
        },
        "willConnectAutomatically": " will connect automatically."
      }
    },
    "error": {
      "forbidden": "Forbidden.",
      "unauthorized": "Unauthorized.",
      "notFound": "User is not found.",
      "internalServerError": "User error occured."
    }
  },
  "oauth": {
    "error": {
      "notCreated": "Oauth connection is not created.",
      "forbidden": "Forbidden.",
      "notFound": "Oauth connection is not found.",
      "internalServerError": "Oauth connection error."
    }
  },
  "profile": {
    "title": "Profile",
    "follow": {
      "title": "Follow",
      "done": "You are following %s now.",
      "fail": "Sorry, following perform failed."
    },
    "unfollow": {
      "title": "Unfollow",
      "done": "You are unfollowing %s now.",
      "fail": "Sorry, unfollowing perform failed."
    },
    "following": {"title": "Following", "more": "more Following"},
    "followers": {"title": "Followers", "more": "more Followers"}
  },
  "setting": {
    "basic": {"title": "Basic"},
    "password": {
      "title": "Change Password",
      "currentPassword": "Current Password",
      "newPassword": "New Password",
      "newPasswordConfirm": "New Password Confirm",
      "enterCurrentPassword": "Please enter your current password.",
      "enterNewPassword": "Please enter new password.",
      "enterNewPasswordConfirm": "Please enter your new password again.",
      "done": "User password is updated.",
      "fail": "User password is not updated.",
      "passwordIncorrect":
          "You entered incorrect password. Please enter correct password.",
      "submit": "Save changes"
    },
    "connection": {"title": "Connection"},
    "leaveOurService": {
      "title": "Leave our service",
      "thankYou": [
        "We were happy with you.",
        "If you want to come back later,",
        "You are more than welcome.",
        "Thank for use our service."
      ],
      "done": "Bye Bye. We hope to see you again.",
      "fail": "Account deletion failed.",
      "confirm": {
        "title": "Are you sure to leave?",
        "content":
            "If you decide to leave our service, your account will be delete permenantly.",
        "close": "Close",
        "submit": "I'm sure to leave"
      },
      "wantToLeave": "I want to leave here"
    }
  },
  "article": {
    "modelName": "article",
    "title": "Title",
    "content": "Content",
    "category": {"notice": "Notice", "general": "General", "etc": "ETC"},
    "view": {
      "tags": "Tags",
      "form": {
        "list": "List",
        "item": "Item",
        "category": "Category",
        "new": {"title": "Write an article", "submit": "Create"},
        "edit": {"title": "Edit an article", "submit": "Update"},
        "placeholder": {
          "title": "My article",
          "content": "Content",
          "tags": "Comma Separted tags. EX) health,coffee,hand drip"
        },
        "validation": {
          "title": "Please enter a title.",
          "content": "Please enter a content.",
          "tags": "Please enter tags."
        }
      },
      "list": {"title": "Life", "write": "Write", "more": "More"},
      "item": {
        "list": "List",
        "edit": "Edit",
        "delete": "Delete",
        "confirm": {
          "title": "Delete Post",
          "content": "Do you really want to delete this article?",
          "cancelText": "Cancel",
          "confirmText": "Delete"
        },
        "author": {"fan": "fan"}
      },
      "created": {
        "done": "Article is created successfully.",
        "fail": "Article creation is failed."
      },
      "updated": {
        "title": "Update article",
        "done": "Article is successfully updated.",
        "fail": "Article update is failed."
      },
      "deleted": {
        "done": "Article is successfully deleted.",
        "fail": "Article deletion is failed."
      }
    },
    "error": {
      "notCreated": "Article is not created.",
      "isNotAuthor": "Author only can modify the article.",
      "forbidden": "Forbidden.",
      "notFound": "Article is not found."
    }
  },
  "comment": {
    "list": {
      "count_zero": "Leave a first comment.",
      "count_many": "{{count}} comment",
      "count_many_plural": "{{count}} comments"
    },
    "created": {
      "done": "The comment is saved.",
      "fail": "The comment is not saved."
    },
    "updated": {
      "done": "The comment updated successfully.",
      "fail": "The comment did not updated."
    },
    "placeholder": {"content": "Please leave a comment."},
    "confirmDelete": {
      "title": "Are you sure you want to delete the comment?",
      "content":
          "If you confirm to delete the comment, it's going to be removed.",
      "cancelText": "Cancel",
      "confirmText": "Delete"
    },
    "deleted": {
      "done": "The comment is deleted.",
      "fail": "The comment is not deleted.",
      "conflict": "The comment can't be deleted because of the replies."
    },
    "error": {
      "unauthorized": "Unauthorized.",
      "forbidden": "User has no permission.",
      "notFound": "Comment or parent is not found."
    },
    "form": {
      "new": {"submit": "Create a comment"},
      "edit": {"submit": "Update a comment"},
      "placeholder": "Put your comment"
    },
    "showMore": "Show more comments"
  },
  "liking": {
    "like": {
      "title": "Like",
      "done": "You like this item.",
      "fail": "Sorry, liked failed."
    },
    "unlike": {
      "title": "Unlike",
      "done": "You unlike this item.",
      "fail": "Sorry, unliked failed."
    },
    "likings": {"title": "Likings", "more": "more Likings"},
    "fan": {
      "title": "FAN",
      "follow": "+",
      "unFollow": "-",
      "more": "Show more fan."
    },
    "error": {
      "unauthorized": "Unauthorized.",
      "forbidden": "User has no permission.",
      "notFound": "Object not found."
    }
  },
  "upload": {
    "done": "Files are uploaded successfully.",
    "error": {
      "unauthorized": "Unauthorized.",
      "internalServerError": "Files are not uploaded."
    },
    "file": {
      "create": {
        "done": "File meta created successfully.",
        "fail": "File meta not created successfully."
      },
      "update": {
        "done": "File meta updated successfully.",
        "fail": "File meta not updated successfully."
      },
      "delete": {
        "done": "File meta deleted successfully.",
        "fail": "File meta not deleted successfully."
      },
      "error": {
        "unauthorized": "Unauthorized.",
        "notFound": "File meta not found."
      }
    }
  },
  "facebook": {
    "title": "Facebook",
    "sendMessage": "Send Message To Facebook",
    "messagePlaceholder": "It's a great service!!",
    "connect": "Connect to Facebook",
    "disconnect": "Disconnect to Facebook",
    "send": {
      "done": "Message sent.",
      "fail": "Sorry, message sending failed.",
      "connectionFailed":
          "Sorry, connection failed. Please connect to facebook."
    }
  },
  "sharing": {
    "facebook": {
      "title": "Share on Facebook",
      "done": "The article shared into your facebook wall."
    }
  },
  "howItWorks": {
    "title": "How it works",
    "content": ["%(name)s works well.", ""]
  }
};
