<?xml version="1.0" encoding="UTF-8"?>
<tileset version="1.2" tiledversion="1.2.1" name="monday5" tilewidth="16" tileheight="16" tilecount="440" columns="22">
 <image source="../platformer/assets/tilesheets/platformer2.png" width="352" height="320"/>
 <terraintypes>
  <terrain name="New Terrain" tile="79"/>
 </terraintypes>
 <tile id="6" type="Enemy">
  <objectgroup draworder="index">
   <object id="2" x="4" y="4" width="8" height="12"/>
  </objectgroup>
 </tile>
 <tile id="35">
  <objectgroup draworder="index">
   <object id="1" x="4" y="0">
    <polygon points="0,0 0,8 4,8 4,12 8,16 12,16 12,0"/>
   </object>
   <object id="2" x="4" y="8">
    <polygon points="0,0 -4,4 0,4"/>
   </object>
   <object id="3" x="4" y="8">
    <polygon points="0,0 -4,4 0,4"/>
   </object>
  </objectgroup>
 </tile>
 <tile id="79">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="80">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="81">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="101">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="102">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="103">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="down" type="bool" value="true"/>
     <property name="up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="185" type="block">
  <properties>
   <property name="hakan" value="test"/>
   <property name="magnus" type="bool" value="false"/>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="186" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" name="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="187" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="3" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="188" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="207" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="208" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="209" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="5" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="210" type="block">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
</tileset>
