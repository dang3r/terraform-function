terraform-provider-function
=======================

A terraform plugin that uses terraform resources to help execute functions on data that are not easily supported by builtin-functions.

## Building
    go get github.com/dang3r/terraform-function
    cd $GO_WORKSPACE/src/github.com/danger/terraform-function
    make 


## Configuring
Terraform needs to detect the plugin binary. This can be done in two ways
 * Copy the ```terraform-provider-function``` binary into the root directory of your terraform project
 * Add the following to your ```~/.terraformrc```
 ```json
providers {
    ...
    function = "$GOLANG_WORKSPACE/src/github.com/dang3r/terraform-function/terraform-provider-function"
}
```
## Examples
View the example folder to see sample use. Note that after a resource representing a function is executed, the return value can be retrieved from the `return` attribute of the resource (`foo.return`)

## Supported Functions
* map_from_list
  * Given a list of [key1, val1, key2, val2, ...], convert it into a map in the form 
    ```json
    {  
        key1 : val1,
        key2 : val2,
        ...
    }
    ```
* zip
  *  Given two lists in the form [key1, key2, key3,...], [val1, val2, val3, ...] convert it into a map in the form
        ```json
        {  
            key1 : val1,
            key2 : val2,
            ...
        }
        ```