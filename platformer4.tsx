<?xml version="1.0" encoding="UTF-8"?>
<tileset version="1.2" tiledversion="1.2.1" name="monday5" tilewidth="16" tileheight="16" tilecount="520" columns="26">
 <image source="../platformer/assets/tilesheets/platformer2.png" width="416" height="320"/>
 <terraintypes>
  <terrain name="New Terrain" tile="79"/>
 </terraintypes>
 <tile id="6" type="Enemy">
  <properties>
   <property name="bouncy" type="bool" value="false"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" x="4" y="4" width="8" height="12"/>
  </objectgroup>
 </tile>
 <tile id="7">
  <properties>
   <property name="bouncy" type="bool" value="false"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" x="4" y="6" width="8" height="10"/>
  </objectgroup>
 </tile>
 <tile id="8">
  <properties>
   <property name="bouncy" type="bool" value="false"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
 </tile>
 <tile id="9">
  <properties>
   <property name="bouncy" type="bool" value="true"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="12" x="2" y="4" width="12" height="12"/>
  </objectgroup>
 </tile>
 <tile id="10">
  <properties>
   <property name="bouncy" type="bool" value="true"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
 </tile>
 <tile id="11">
  <properties>
   <property name="bouncy" type="bool" value="true"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
 </tile>
 <tile id="12">
  <properties>
   <property name="bouncy" type="bool" value="true"/>
   <property name="hazard" type="bool" value="true"/>
   <property name="killable" type="bool" value="true"/>
  </properties>
 </tile>
 <tile id="20">
  <objectgroup draworder="index">
   <object id="2" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="21">
  <objectgroup draworder="index">
   <object id="2" x="0" y="8" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="22">
  <objectgroup draworder="index">
   <object id="4" x="0" y="0" width="8" height="16"/>
  </objectgroup>
 </tile>
 <tile id="23">
  <objectgroup draworder="index">
   <object id="1" x="8" y="0" width="8" height="16"/>
  </objectgroup>
 </tile>
 <tile id="24">
  <objectgroup draworder="index">
   <object id="1" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="39">
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
 <tile id="46">
  <objectgroup draworder="index">
   <object id="3" x="0" y="5" width="16" height="6"/>
  </objectgroup>
 </tile>
 <tile id="47">
  <objectgroup draworder="index">
   <object id="2" x="5" y="0" width="6" height="16"/>
  </objectgroup>
 </tile>
 <tile id="48">
  <objectgroup draworder="index">
   <object id="2" x="0" y="7" width="16" height="2"/>
  </objectgroup>
 </tile>
 <tile id="49">
  <objectgroup draworder="index">
   <object id="5" x="7" y="0" width="2" height="16"/>
  </objectgroup>
 </tile>
 <tile id="72">
  <objectgroup draworder="index">
   <object id="5" x="0" y="0" width="16" height="2"/>
  </objectgroup>
 </tile>
 <tile id="73">
  <objectgroup draworder="index">
   <object id="1" x="0" y="14" width="16" height="2"/>
  </objectgroup>
 </tile>
 <tile id="74">
  <objectgroup draworder="index">
   <object id="1" x="0" y="0" width="2" height="16"/>
  </objectgroup>
 </tile>
 <tile id="75">
  <objectgroup draworder="index">
   <object id="1" x="14" y="0" width="2" height="16"/>
  </objectgroup>
 </tile>
 <tile id="91">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="92">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="93">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="117">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="118">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="119">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8">
    <properties>
     <property name="allow_from_down" type="bool" value="true"/>
     <property name="allow_from_up" type="bool" value="false"/>
    </properties>
   </object>
  </objectgroup>
 </tile>
 <tile id="125" type=" ">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="8" width="16" height="4"/>
  </objectgroup>
 </tile>
 <tile id="126" type=" ">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="8" width="16" height="4"/>
  </objectgroup>
 </tile>
 <tile id="127" type=" ">
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="8" width="16" height="4"/>
  </objectgroup>
 </tile>
 <tile id="185" type="block"/>
 <tile id="186" type="block"/>
 <tile id="187" type="block"/>
 <tile id="188" type="block"/>
 <tile id="207" type="block"/>
 <tile id="208" type="block"/>
 <tile id="209" type="block"/>
 <tile id="210" type="block"/>
 <tile id="217">
  <properties>
   <property name="hakan" value="test"/>
   <property name="magnus" type="bool" value="false"/>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="218">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" name="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="219">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="3" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="220">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="2" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="243">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="244">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
 <tile id="245">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="5" type="hitbox" x="0" y="0" width="16" height="16"/>
  </objectgroup>
 </tile>
 <tile id="246">
  <properties>
   <property name="solid" type="bool" value="true"/>
  </properties>
  <objectgroup draworder="index">
   <object id="1" type="hitbox" x="0" y="0" width="16" height="8"/>
  </objectgroup>
 </tile>
</tileset>
