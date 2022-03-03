(provide 'get-task-details)

(defun get-task-details (jlink)
  (substring
   (shell-command-to-string (format "/home/user/programming/personal/ejira/emacs-go --operation OrgJiraDetails --value %s" jlink))
   9 -2))

