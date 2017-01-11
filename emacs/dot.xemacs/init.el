;; (set-background-color "DarkSlateGray")
;; (set-foreground-color "white")
;; (set-cursor-color "red")

;; (tool-bar-mode 0)
;; (set-scroll-bar-mode nil)
;; (defun set-colors ()
;;    "Set fixed colors"
;;    (interactive)
;;    (set-background-color "DarkSlateGray")
;;    (set-foreground-color "white")
;;    (set-cursor-color "red")
;; )
;; (set-colors)

;; (set-default-font "8x13")
;; (set-face-attribute 'default nil :font "Monospace-12")

(autoload 'c++-mode "cc-mode"  "C++ Editing Mode" t)
(autoload 'c-mode "cc-mode" "C Editing Mode" t)

(setq auto-mode-alist
    (append '(("\\.C$" . c++-mode)
        ("\\.cc$" . c++-mode)
        ("\\.cpp$" . c++-mode)
        ("\\.cxx$" . c++-mode)
        ("\\.hh$" . c++-mode)
        ("\\.hpp$" . c++-mode)
        ("\\.h$" . c++-mode)
        ("\\.c$" . c-mode)
        ) auto-mode-alist)
)

(setq-default fill-column 80)
(setq auto-fill-mode 1)
(setq gdb-many-windows 1)
(global-set-key '[f1] 'other-window)
(global-set-key '[f2] 'gud-up)
(global-set-key '[f3] 'gud-down)
(global-set-key '[f4] 'c-comment)
(global-set-key '[f5] 'gud-break)
(global-set-key '[f6] 'erase-buffer)
(global-set-key '[f7] 'revert-buffer)
(global-set-key '[f8] 'goto-line)
(global-set-key '[f9] 'shell)
(global-set-key '[f10] 'gud-next)
(global-set-key '[f11] 'gud-step)
(global-set-key '[f12] 'undo)
;; (global-set-key '[M-left] 'backward-sexp)
;; (global-set-key '[M-right] 'forward-sexp)



(setq c-mode-hook
    (function (lambda ()
    (setq indent-tabs-mode nil)
    (setq c-indent-level 4))))

(setq objc-mode-hook
(function (lambda ()
            (setq indent-tabs-mode nil)
            (setq c-indent-level 4))))

(setq c++-mode-hook
      (function (lambda ()
                  (setq indent-tabs-mode nil)
                  (setq c-indent-level 4))))


(setq inhibit-startup-message t)
(setq require-final-newline t)
(setq display-time-day-and-date t)
(setq search-highlight t)
(setq next-line-add-newlines nil); down key wont add newline
(setq indent-tabs-mode nil)
(setq default-tab-width 2)
(setq bell-volume 0)

;; Herman c++ indent function. Primarily for namespaces
(defun my-c++-setup ()
    (setq c-basic-offset 2)
    (setq indent-tabs-mode nil)
    (c-set-offset 'innamespace 0)
    (c-set-offset 'arglist-intro '++)
    (c-set-offset 'arglist-cont 0)
    (c-set-offset 'arglist-close 0)
    (c-set-offset 'substatement-open 0)
)

;; (gud-tooltip-mode 1)

(add-hook 'c++-mode-hook 'my-c++-setup)
;; (set-background-color "DarkSlateGray")
;; (set-foreground-color "white")
;; (set-cursor-color "red")

;; (defun set-colors ()
;;    "Set fixed colors"
;;    (interactive)
;;    (set-background-color "DarkSlateGray")
;;   (set-foreground-color "white")
;;    (set-cursor-color "red")
;; )
;; (set-colors)

;; (set-default-font "8x13")
;; (set-face-attribute 'default nil :font "Monospace-12")

(autoload 'c++-mode "cc-mode"  "C++ Editing Mode" t)
(autoload 'c-mode "cc-mode" "C Editing Mode" t)

(setq-default fill-column 80)
(setq auto-fill-mode 1)

(setq gdb-many-windows 1)
(global-set-key '[f1] 'other-window)
(global-set-key '[f2] 'gud-up)
(global-set-key '[f3] 'gud-down)
(global-set-key '[f4] 'c-comment)
(global-set-key '[f5] 'gud-break)
(global-set-key '[f6] 'erase-buffer)
(global-set-key '[f7] 'revert-buffer)
(global-set-key '[f8] 'goto-line)
(global-set-key '[f9] 'shell)
(global-set-key '[f10] 'gud-next)
(global-set-key '[f11] 'gud-step)
(global-set-key '[f12] 'undo)
;; (global-set-key '[M-left] 'backward-sexp)
;; (global-set-key '[M-right] 'forward-sexp)




(setq inhibit-startup-message t)
(setq require-final-newline t)
(setq display-time-day-and-date t)
(setq search-highlight t)
(setq next-line-add-newlines nil); down key wont add newline
(setq indent-tabs-mode nil)
(setq default-tab-width 2)
(setq bell-volume 0)

;; (gud-tooltip-mode 1)

;; Herman c++ indent function. Primarily for namespaces
(defun my-c++-setup ()
    (setq c-basic-offset 4)
    (setq indent-tabs-mode nil)
    (c-set-offset 'innamespace 0)
    (c-set-offset 'arglist-intro '++)
    (c-set-offset 'arglist-cont 0)
    (c-set-offset 'arglist-close 0)
    (c-set-offset 'substatement-open 0)
)

(add-hook 'c++-mode-hook 'my-c++-setup)

(add-hook 'c++-mode-hook 'my-c++-setup)

(global-auto-revert-mode t)

;;pdb setup, note the python version
(setq pdb-path '/usr/lib64/python2.7/pdb.py
    gud-pdb-command-name (symbol-name pdb-path))

;; X (setq debug-on-error 't)

;; define macro for inserting comment
(fset 'c-comment
    "/**\C-m\C-i*\C-m*/\C-[OA ")
;; define macro for inserting header
(fset 'c-header
    "//------------------------------------------------------------------------------\C-m")



(setq line-number-mode t)
(line-number-mode t)
;; (setq column-number-mode t)

;;=======================================================================
;; in my ~/.Xresources and ~/.emacs, respectively.
;; or in ~/.xemacs/init.el or ~/.xemacs/custom.el
;; Emacs.font: Liberation Mono-11

;; (set-default-font "Liberation Mono-11")
(set-face-background 'default  "DarkSlateGray")
(set-face-foreground 'default  "white")

;;=======================================================================
;; This is for debugging (gdb in xemacs)
;;=======================================================================
(set-face-background 'gdb-arrow-face  "Blue")
;; (set-face-foreground 'gdb-arrow-face  "white")

(set-default-font "Font -*-Liberation Mono-medium-r-*-*-*-120-*-*-*-*-iso8859-*")

