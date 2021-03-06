(provide 'get-task-details)

(defun get-task-details (jlink)
  (substring
   (shell-command-to-string (format "/home/user/programming/personal/ejira/emacs-go --operation OrgJiraDetails --value %s" jlink))
   0 -1))

(defun get-jira-id-from-buffer()
  (goto-char (point-min))
  (re-search-forward "Link: .*?browse/\\(.*?\\)\]")
  (match-string-no-properties 1)
  )

(defun update-jira-information ()
  (interactive)
  (setq jlink (get-jira-id-from-buffer))
  (goto-char (point-min))
  (search-forward "* Jira Information")
  (org-mark-subtree)
  (delete-backward-char 1)
  (ufg-add-jira-link-information jlink)
)

(defun put-day-to-jira-issue ()
  (interactive)
  (org-copy-subtree 1)
  (setq file (make-temp-file "ejira.tmp"))
  (with-temp-file file
    (org-paste-subtree)
    )
  (save-excursion
    (setq jlink (get-jira-id-from-buffer))
    (shell-command-to-string (format "/home/user/programming/personal/ejira/emacs-go --operation AddComment --value %s --vfile %s" jlink file))
    (delete-file file)))

(defun put-to-jira-issue ()
  (interactive)
  (setq project (read-string "Enter a Project ID (or name) to send this to: "))
  (goto-char (point-min))
  (set-mark-command nil)
  (org-next-visible-heading 4)
  (kill-ring-save (region-beginning) (region-end))
  (setq file (make-temp-file "ejira.tmp"))
  (with-temp-file file
    (yank)
    )
  (setq issue
	(shell-command-to-string (format "/home/user/programming/personal/ejira/emacs-go --operation AddIssue --value %s --vfile %s 2>/dev/null" project file))
	)
  (org-return)
  (previous-line)
  (ufg-add-jira-link-information issue)
  (delete-file file))


