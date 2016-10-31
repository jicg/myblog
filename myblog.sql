/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50715
 Source Host           : localhost
 Source Database       : myblog

 Target Server Type    : MySQL
 Target Server Version : 50715
 File Encoding         : utf-8

 Date: 09/25/2016 23:04:40 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `jicg_article`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_article`;
CREATE TABLE `jicg_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `content` text,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `art_type` int(11) DEFAULT NULL,
  `source` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_jicg_article_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `jicg_article`
-- ----------------------------
BEGIN;
INSERT INTO `jicg_article` VALUES ('3', '版权声明', '版权声明', '<h3>版权声明</h3>\r\n<p>1.本网站的版权归原作者所有。<br />\r\n2.本站仅供学习与参考。<br />\r\n3.本站在建设中引用了因特网上的一些资源。若有内容侵犯了您的权益，请联系我们，我们可做调整。<br />\r\n4.转载本站原创资源，请注明来自本站。<br />\r\n5.如果您对网站有建设性的好意见，请联系我们。</p>', '2016-07-07 10:17:43', '2016-07-07 10:17:43', '1', null), ('4', '宣传页面版本说名', '宣传页面版本说名', '<h2 style=\"text-align: center;\">宣传声明</h2>\r\n\r\n<p style=\"margin-left: 40px;\"><span style=\"font-size:16px\">1、本模块内的文章 是用来做宣传的用的，</span></p>\r\n\r\n<p style=\"margin-left: 40px;\"><span style=\"font-size:16px\">2、本模块内的文章，如果侵犯他人的公民权力，请联系站长。</span></p>\r\n\r\n<hr />\r\n<p style=\"text-align: center;\">edit by jicg&nbsp;</p>\r\n', '2016-07-07 20:00:23', '2016-07-07 20:00:23', '1', null), ('7', '你好世界', '你好世界，欢迎来到golang世界', '<h2 id=\"h2--\"><a name=\"你好，世界\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>你好，世界</h2><pre><code class=\"lang-go\">package admin\r\n\r\nimport(\r\n    \"fmt\"\r\n)\r\n\r\nfunc main(){\r\n    fmt.print(\"hello world!!\")\r\n}\r\n</code></pre>\r\n', '2016-07-10 10:22:03', '2016-07-11 09:15:33', '2', '## 你好，世界\r\n\r\n```go\r\npackage admin\r\n\r\nimport(\r\n	\"fmt\"\r\n)\r\n\r\nfunc main(){\r\n	fmt.print(\"hello world!!\")\r\n}\r\n\r\n```\r\n\r\n\r\n'), ('12', 'oracle 临时表 资源占用', 'Oracle 资源占用', '<h3 id=\"h3-u6743u9650\"><a name=\"权限\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>权限</h3><pre><code class=\"lang-sql\">grant select any dictionary to XXX\r\ngrant alter system to xxx\r\n</code></pre>\r\n<h3 id=\"h3--sid\"><a name=\"查询sid\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>查询sid</h3><pre><code>select sid, serial#\r\n  from v$session\r\n where sid = (select sid\r\n    from v$lock\r\n    where id1 = (\r\n        select object_id\r\n        from user_objects\r\n        where object_name =upper(&#39;tablename&#39;)));\r\n</code></pre><h3 id=\"h3-kill-sid\"><a name=\"Kill sid\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>Kill sid</h3><pre><code>alter system kill session &#39;sid,serial#&#39;;\r\n</code></pre>', '2016-08-09 21:20:26', '2016-08-09 21:20:26', '2', '###权限\r\n```sql\r\ngrant select any dictionary to XXX\r\ngrant alter system to xxx \r\n```\r\n\r\n###查询sid\r\n```\r\nselect sid, serial#\r\n  from v$session\r\n where sid = (select sid\r\n    from v$lock\r\n    where id1 = (\r\n		select object_id\r\n        from user_objects\r\n        where object_name =upper(\'tablename\')));\r\n```\r\n\r\n###Kill sid\r\n```\r\nalter system kill session \'sid,serial#\';\r\n```'), ('13', 'Android', 'Android代码采集（一）', '<h6 id=\"h6-android-\"><a name=\"android 让标题栏下拉按钮显示\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>android 让标题栏下拉按钮显示</h6><pre><code class=\"lang-java\">  protected void onCreate(Bundle savedInstanceState){\r\n        super.onCreate(savedInstanceState);\r\n        setOverflowShowingAlways();\r\n  }\r\n\r\nprivate void setOverflowShowingAlways() {\r\n        try {\r\n            ViewConfiguration config = ViewConfiguration.get(this);\r\n            Field menuKeyField = ViewConfiguration.class.getDeclaredField(&quot;sHasPermanentMenuKey&quot;);\r\n            menuKeyField.setAccessible(true);\r\n            menuKeyField.setBoolean(config, false);\r\n        } catch (Exception e) {\r\n            e.printStackTrace();\r\n        }\r\n    }\r\n</code></pre>\r\n<h6 id=\"h6-android-\"><a name=\"android 下拉菜单中的图片显示\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>android 下拉菜单中的图片显示</h6><pre><code class=\"lang-java\">    @Override\r\n    public boolean onMenuOpened(int featureId, Menu menu) {\r\n        if (featureId == Window.FEATURE_ACTION_BAR &amp;&amp; menu != null) {\r\n            if (menu.getClass().getSimpleName().equals(&quot;MenuBuilder&quot;)) {\r\n                try {\r\n                    Method m = menu.getClass().getDeclaredMethod(&quot;setOptionalIconsVisible&quot;, Boolean.TYPE);\r\n                    m.setAccessible(true);\r\n                    m.invoke(menu, true);\r\n                } catch (Exception e) {\r\n                }\r\n            }\r\n        }\r\n        return super.onMenuOpened(featureId, menu);\r\n    }\r\n</code></pre>\r\n<h5 id=\"h5-android-sharedpreferences-\"><a name=\"android SharedPreferences的使用\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>android SharedPreferences的使用</h5><pre><code class=\"lang-java\">SharedPreferences share = context.getSharedPreferences(&quot;perference&quot;, Context.MODE_PRIVATE);\r\n// 取出pwd\r\nString value = share.getString(&quot;pwd&quot;, &quot;123&quot;)\r\nSharedPreferences.Editor editor = share.edit();\r\n// 保存pwd\r\neditor.putString(&quot;pwd&quot;, &quot;123&quot;);\r\n</code></pre>\r\n', '2016-08-22 17:44:03', '2016-08-22 18:02:50', '2', '######android 让标题栏下拉按钮显示\r\n```java\r\n  protected void onCreate(Bundle savedInstanceState){\r\n        super.onCreate(savedInstanceState);\r\n        setOverflowShowingAlways();\r\n  }\r\n\r\nprivate void setOverflowShowingAlways() {\r\n        try {\r\n            ViewConfiguration config = ViewConfiguration.get(this);\r\n            Field menuKeyField = ViewConfiguration.class.getDeclaredField(\"sHasPermanentMenuKey\");\r\n            menuKeyField.setAccessible(true);\r\n            menuKeyField.setBoolean(config, false);\r\n        } catch (Exception e) {\r\n            e.printStackTrace();\r\n        }\r\n    }\r\n```\r\n\r\n###### android 下拉菜单中的图片显示\r\n```java\r\n	@Override\r\n    public boolean onMenuOpened(int featureId, Menu menu) {\r\n        if (featureId == Window.FEATURE_ACTION_BAR && menu != null) {\r\n            if (menu.getClass().getSimpleName().equals(\"MenuBuilder\")) {\r\n                try {\r\n                    Method m = menu.getClass().getDeclaredMethod(\"setOptionalIconsVisible\", Boolean.TYPE);\r\n                    m.setAccessible(true);\r\n                    m.invoke(menu, true);\r\n                } catch (Exception e) {\r\n                }\r\n            }\r\n        }\r\n        return super.onMenuOpened(featureId, menu);\r\n    }\r\n```\r\n#####android SharedPreferences的使用\r\n```java\r\nSharedPreferences share = context.getSharedPreferences(\"perference\", Context.MODE_PRIVATE);\r\n// 取出pwd\r\nString value = share.getString(\"pwd\", \"123\")\r\nSharedPreferences.Editor editor = share.edit();\r\n// 保存pwd\r\neditor.putString(\"pwd\", \"123\");\r\n```\r\n\r\n');
COMMIT;

-- ----------------------------
--  Table structure for `jicg_blog`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_blog`;
CREATE TABLE `jicg_blog` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `remark` text,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `jicg_category`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_category`;
CREATE TABLE `jicg_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `orderno` bigint(20) DEFAULT NULL,
  `pid` bigint(20) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `article` text,
  `art_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_jicg_category_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `jicg_category`
-- ----------------------------
BEGIN;
INSERT INTO `jicg_category` VALUES ('1', 'go成长记录', 'go语言的感悟，语法，以及大神的代卖', '10', '0', '2016-07-04 21:05:43', '2016-07-07 17:38:26', '123', '3'), ('2', '语法熟悉（一）', '熟悉语法', '20', '1', '2016-07-04 21:34:24', '2016-07-10 12:21:19', '', '7'), ('13', 'Oracle成长记录', '', '200', '0', '2016-08-09 21:21:19', '2016-08-09 21:21:34', '', '3'), ('14', 'Oracle资源占用', 'Oracle临时表占用', '10', '13', '2016-08-09 21:22:06', '2016-08-09 21:22:25', '', '12'), ('15', 'Android', '安卓的总结', '30', '0', '2016-08-22 17:22:06', '2016-08-22 17:44:17', '', '13');
COMMIT;

-- ----------------------------
--  Table structure for `jicg_ideal`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_ideal`;
CREATE TABLE `jicg_ideal` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `remark` text,
  `cdate` datetime DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `jicg_image`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_image`;
CREATE TABLE `jicg_image` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `file_path` varchar(255) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `file_size` bigint(20) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `file_type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_jicg_image_name` (`name`),
  UNIQUE KEY `UQE_jicg_image_uri` (`uri`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `jicg_nav`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_nav`;
CREATE TABLE `jicg_nav` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orderno` int(11) DEFAULT '0',
  `name` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `issys` tinyint(1) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `delete` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_jicg_nav_url` (`url`),
  KEY `IDX_jicg_nav_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `jicg_nav`
-- ----------------------------
BEGIN;
INSERT INTO `jicg_nav` VALUES ('1', '1', '基础信息', null, '/admin', '0', '0', '1', '2016-04-10 20:00:59', '2016-04-10 20:00:59', null), ('2', '30', '系统设置', null, '/admin/sys', '0', '1', '1', '2016-04-10 20:01:53', '2016-04-10 20:01:53', null), ('3', '20', '导航设置', null, '/admin/sys/navs/list', '2', '1', '1', '2016-04-10 20:02:44', '2016-04-10 20:02:44', null), ('4', '10', '参数设置', null, '/admin/sys/param/list', '2', '0', '1', '2016-04-10 20:04:37', '2016-07-03 13:51:53', '2016-07-16 18:26:06'), ('5', '2', '测试', null, '/admin/index', '1', '0', '1', '2016-04-11 16:15:40', '2016-04-11 16:15:40', '2016-07-16 18:26:02'), ('6', '50', '文章', '', '/admin/article/list', '1', '0', '1', '2016-07-03 11:21:32', '2016-07-06 16:29:19', null), ('7', '10', '目录', '', '/admin/category/list', '1', '0', '1', '2016-07-03 11:22:56', '2016-07-05 21:57:59', null), ('9', '90', '图片管理', '', '/admin/image/list', '1', '0', '1', '2016-07-03 21:46:22', '2016-07-03 21:46:22', null);
COMMIT;

-- ----------------------------
--  Table structure for `jicg_user`
-- ----------------------------
DROP TABLE IF EXISTS `jicg_user`;
CREATE TABLE `jicg_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `issys` tinyint(1) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `is_make_down` tinyint(1) DEFAULT NULL,
  `editor_type` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_jicg_user_name` (`name`),
  UNIQUE KEY `UQE_jicg_user_email` (`email`),
  KEY `IDX_jicg_user_nickname` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `jicg_user`
-- ----------------------------
BEGIN;
INSERT INTO `jicg_user` VALUES ('1', 'jicg', 'jicg', 'ji.cg@qq.com', 'jicg', '1', '1', '2016-07-03 16:08:23', '2016-07-03 16:08:31', null, '1'), ('2', 'admin', 'admin', 'admin@qq.com', 'admin', '1', '1', null, null, null, '2');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
