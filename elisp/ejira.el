(provide 'get-task-details)

(defun get-task-details (jlink)
  (substring
   (shell-command-to-string (format "/home/user/programming/personal/ejira/emacs-go --operation OrgJiraDetails --value %s" jlink))
   9 -2))

(defun update-jira-information ()
  (interactive)
  (goto-char (point-min))
  (re-search-forward "Link: .*?browse/\\(.*?\\)\]")
  (setq jlink (match-string-no-properties 1))
  (goto-char (point-min))
  (search-forward "* Jira Information")
  (org-mark-subtree)
  (delete-backward-char 1)
  (ufg-add-jira-link-information jlink)
)
