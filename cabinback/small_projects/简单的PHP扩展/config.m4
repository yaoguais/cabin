dnl $Id$
dnl config.m4 for extension democlass

dnl Comments in this file start with the string 'dnl'.
dnl Remove where necessary. This file will not work
dnl without editing.

dnl If your extension references something external, use with:

PHP_ARG_WITH(democlass, for democlass support,
Make sure that the comment is aligned:
[  --with-democlass             Include democlass support])

dnl Otherwise use enable:

dnl PHP_ARG_ENABLE(democlass, whether to enable democlass support,
dnl Make sure that the comment is aligned:
dnl [  --enable-democlass           Enable democlass support])

if test "$PHP_DEMOCLASS" != "no"; then
  dnl Write more examples of tests here...

  dnl # --with-democlass -> check with-path
  dnl SEARCH_PATH="/usr/local /usr"     # you might want to change this
  dnl SEARCH_FOR="/include/democlass.h"  # you most likely want to change this
  dnl if test -r $PHP_DEMOCLASS/$SEARCH_FOR; then # path given as parameter
  dnl   DEMOCLASS_DIR=$PHP_DEMOCLASS
  dnl else # search default path list
  dnl   AC_MSG_CHECKING([for democlass files in default path])
  dnl   for i in $SEARCH_PATH ; do
  dnl     if test -r $i/$SEARCH_FOR; then
  dnl       DEMOCLASS_DIR=$i
  dnl       AC_MSG_RESULT(found in $i)
  dnl     fi
  dnl   done
  dnl fi
  dnl
  dnl if test -z "$DEMOCLASS_DIR"; then
  dnl   AC_MSG_RESULT([not found])
  dnl   AC_MSG_ERROR([Please reinstall the democlass distribution])
  dnl fi

  dnl # --with-democlass -> add include path
  dnl PHP_ADD_INCLUDE($DEMOCLASS_DIR/include)

  dnl # --with-democlass -> check for lib and symbol presence
  dnl LIBNAME=democlass # you may want to change this
  dnl LIBSYMBOL=democlass # you most likely want to change this 

  dnl PHP_CHECK_LIBRARY($LIBNAME,$LIBSYMBOL,
  dnl [
  dnl   PHP_ADD_LIBRARY_WITH_PATH($LIBNAME, $DEMOCLASS_DIR/$PHP_LIBDIR, DEMOCLASS_SHARED_LIBADD)
  dnl   AC_DEFINE(HAVE_DEMOCLASSLIB,1,[ ])
  dnl ],[
  dnl   AC_MSG_ERROR([wrong democlass lib version or lib not found])
  dnl ],[
  dnl   -L$DEMOCLASS_DIR/$PHP_LIBDIR -lm
  dnl ])
  dnl
  dnl PHP_SUBST(DEMOCLASS_SHARED_LIBADD)

  PHP_NEW_EXTENSION(democlass, democlass.c, $ext_shared,, -DZEND_ENABLE_STATIC_TSRMLS_CACHE=1)
fi
