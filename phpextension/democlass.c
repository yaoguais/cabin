/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) 1997-2015 The PHP Group                                |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

/* $Id$ */

#ifdef HAVE_CONFIG_H
#include "config.h"
#endif

#include "php.h"
#include "php_ini.h"
#include "ext/standard/info.h"
#include "php_democlass.h"

/* If you declare any globals in php_democlass.h uncomment this:
ZEND_DECLARE_MODULE_GLOBALS(democlass)
*/

/* True global resources - no need for thread safety here */
static int le_democlass;

zend_class_entry * democlass_class_entry;

#define DEMOCLASS_ME(func, arg_info, flags) PHP_ME(monster, func, arg_info, flags)
#define DEMOCLASS_MALIAS(func, alias, arg_info, flags) PHP_MALIAS(monster, func, alias, arg_info, flags)
#define DEMOCLASS_METHOD(func) PHP_METHOD(monster, func)

ZEND_BEGIN_ARG_INFO_EX(arginfo_init, 0, 0, 1)
	ZEND_ARG_INFO(0,food)
ZEND_END_ARG_INFO()

ZEND_BEGIN_ARG_INFO_EX(arginfo_eat, 0, 0, 1)
	ZEND_ARG_INFO(0, food)
ZEND_END_ARG_INFO()

ZEND_BEGIN_ARG_INFO_EX(arginfo_add_food, 0, 0, 1)
	ZEND_ARG_INFO(0, food)
ZEND_END_ARG_INFO()

ZEND_BEGIN_ARG_INFO_EX(arginfo_del_food, 0, 0, 1)
	ZEND_ARG_INFO(0, food)
ZEND_END_ARG_INFO()

DEMOCLASS_METHOD(init)
{
	zval * retval;
	if(zend_parse_parameters(ZEND_NUM_ARGS(),"a",&retval) == FAILURE){
		return;
	}
	if(getThis())
	{
		zend_update_property(democlass_class_entry,getThis(),"_foods",sizeof("_foods")-1,retval);
	}
}

DEMOCLASS_METHOD(eat)
{
	int num_idx;
	zval * value, * this_foods, * entry;
	zend_string * str_idx;

	if(zend_parse_parameters(ZEND_NUM_ARGS(),"z",&value) == FAILURE){
		RETURN_FALSE;
	}
	if(!getThis() && Z_TYPE_P(value) != IS_STRING){
		RETURN_FALSE;
	}
	this_foods = zend_read_property(democlass_class_entry,getThis(),"_foods",sizeof("_foods")-1,0,NULL);
	if(Z_TYPE_P(this_foods) != IS_ARRAY){
		RETURN_FALSE;
	}
	ZEND_HASH_FOREACH_KEY_VAL(Z_ARRVAL_P(this_foods), num_idx, str_idx, entry) {
		if (fast_equal_check_string(value, entry)) {
			RETURN_TRUE;
		}
	} ZEND_HASH_FOREACH_END();
	RETURN_FALSE;
}

DEMOCLASS_METHOD(add_food)
{
	zval * value , * this_foods;
	if(zend_parse_parameters(ZEND_NUM_ARGS(),"z",&value) == FAILURE  &&  Z_TYPE_P(value) != IS_STRING && !getThis()){
		return;
	}
	this_foods = zend_read_property(democlass_class_entry,getThis(),"_foods",sizeof("_foods")-1,0,NULL);
	zend_hash_next_index_insert(Z_ARRVAL_P(this_foods),value);
}

DEMOCLASS_METHOD(del_food)
{
	int num_idx;
	zval * value, * this_foods, * entry;
	zend_string * str_idx;
	if(zend_parse_parameters(ZEND_NUM_ARGS(),"z",&value) == FAILURE){
			RETURN_FALSE;
	}
	if(!getThis() && Z_TYPE_P(value) != IS_STRING){
		RETURN_FALSE;
	}
	this_foods = zend_read_property(democlass_class_entry,getThis(),"_foods",sizeof("_foods")-1,0,NULL);
	if(Z_TYPE_P(this_foods) != IS_ARRAY){
		RETURN_FALSE;
	}
	ZEND_HASH_FOREACH_KEY_VAL(Z_ARRVAL_P(this_foods), num_idx, str_idx, entry) {
		if (fast_equal_check_string(value, entry)) {
			zend_hash_index_del(Z_ARRVAL_P(this_foods),num_idx);
		}
	} ZEND_HASH_FOREACH_END();
	RETURN_FALSE;
}

DEMOCLASS_METHOD(get_foods)
{
	zval * retval;
	if(getThis())
	{
		retval = zend_read_property(democlass_class_entry,getThis(),"_foods",sizeof("_foods")-1,0,NULL);
		RETURN_ZVAL(retval,1,NULL);
	}
}

/* {{{ democlass_functions[]
 *
 * Every user visible function must have an entry in democlass_functions[].
 */
const zend_function_entry democlass_functions[] = {
		DEMOCLASS_MALIAS(__construct,	init,	arginfo_init,	ZEND_ACC_PUBLIC)
		DEMOCLASS_ME(eat,arginfo_eat,ZEND_ACC_PUBLIC)
		DEMOCLASS_MALIAS(addFood,add_food,arginfo_add_food,ZEND_ACC_PUBLIC)
		DEMOCLASS_MALIAS(delFood,del_food,arginfo_del_food,ZEND_ACC_PUBLIC)
		DEMOCLASS_MALIAS(getFoods,get_foods,NULL,ZEND_ACC_PUBLIC)
	PHP_FE_END
};
/* }}} */

/* {{{ PHP_MINIT_FUNCTION
 */
PHP_MINIT_FUNCTION(democlass)
{
	/* If you have INI entries, uncomment these lines
	REGISTER_INI_ENTRIES();
	*/

	zend_class_entry ce;

	INIT_CLASS_ENTRY(ce, "monster", democlass_functions);

	democlass_class_entry = zend_register_internal_class_ex(&ce,NULL);

	zend_declare_property_null(democlass_class_entry, "_foods", 		sizeof("_foods") - 1, ZEND_ACC_PRIVATE);

	return SUCCESS;
}
/* }}} */

/* {{{ PHP_MSHUTDOWN_FUNCTION
 */
PHP_MSHUTDOWN_FUNCTION(democlass)
{
	/* uncomment this line if you have INI entries
	UNREGISTER_INI_ENTRIES();
	*/
	return SUCCESS;
}
/* }}} */

/* Remove if there's nothing to do at request start */
/* {{{ PHP_RINIT_FUNCTION
 */
PHP_RINIT_FUNCTION(democlass)
{
#if defined(COMPILE_DL_DEMOCLASS) && defined(ZTS)
	ZEND_TSRMLS_CACHE_UPDATE();
#endif
	return SUCCESS;
}
/* }}} */

/* Remove if there's nothing to do at request end */
/* {{{ PHP_RSHUTDOWN_FUNCTION
 */
PHP_RSHUTDOWN_FUNCTION(democlass)
{
	return SUCCESS;
}
/* }}} */

/* {{{ PHP_MINFO_FUNCTION
 */
PHP_MINFO_FUNCTION(democlass)
{
	php_info_print_table_start();
	php_info_print_table_header(2, "democlass support", "enabled");
	php_info_print_table_end();

	/* Remove comments if you have entries in php.ini
	DISPLAY_INI_ENTRIES();
	*/
}
/* }}} */



/* {{{ democlass_module_entry
 */
zend_module_entry democlass_module_entry = {
	STANDARD_MODULE_HEADER,
	"democlass",
	democlass_functions,
	PHP_MINIT(democlass),
	PHP_MSHUTDOWN(democlass),
	PHP_RINIT(democlass),		/* Replace with NULL if there's nothing to do at request start */
	PHP_RSHUTDOWN(democlass),	/* Replace with NULL if there's nothing to do at request end */
	PHP_MINFO(democlass),
	PHP_DEMOCLASS_VERSION,
	STANDARD_MODULE_PROPERTIES
};
/* }}} */

#ifdef COMPILE_DL_DEMOCLASS
#ifdef ZTS
ZEND_TSRMLS_CACHE_DEFINE();
#endif
ZEND_GET_MODULE(democlass)
#endif

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: noet sw=4 ts=4 fdm=marker
 * vim<600: noet sw=4 ts=4
 */
