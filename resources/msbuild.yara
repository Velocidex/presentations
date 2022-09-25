rule MSBuild_template {
   meta:
      description = "MSBuild template. Detects MSBuild variable setup and generic template strings."
   strings:
      $s1  = "byte[] key_code = new byte[" ascii
      $s2  = "byte[] buff = new byte[" ascii
      $s8  = "<Code Type=\"Class\" Language=\"cs\">" ascii
      $s9  = "<![CDATA[" ascii
      $s10 = "[DllImport(" ascii
      
   condition:
      ( uint16(0) == 0x3c0a or uint8(0) == 0x3c ) // \n< or < at 0
      and any of ($s*) 
}
