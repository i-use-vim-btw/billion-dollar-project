# Gandi Services

A list of all services and services methods.

- Services

  - [Collections](#collections)

  - [Indexes](#indexes)

  - [Vectors](#vectors)

- [All Methods](#all-methods)

## Collections

| Method                | Description                        |
| :-------------------- | :--------------------------------- |
| [create](#create)     | Create a new collection            |
| [create2](#create2)   | Describe a collection              |
| [create3](#create3)   | Drop a collection                  |
| [create4](#create4)   | Get the load state of a collection |
| [create5](#create5)   | Get the stats of a collection      |
| [create6](#create6)   | Check if a collection exists       |
| [create7](#create7)   | List all collections               |
| [create8](#create8)   | Load a collection                  |
| [create9](#create9)   | Release a collection               |
| [create10](#create10) | Rename a collection                |

## Indexes

| Method                | Description        |
| :-------------------- | :----------------- |
| [create11](#create11) | Create a new index |
| [create12](#create12) | Describe an index  |
| [create13](#create13) | Drop an index      |
| [create14](#create14) | List all indexes   |

## Vectors

| Method                | Description    |
| :-------------------- | :------------- |
| [create15](#create15) | Delete vectors |
| [create16](#create16) | Get vectors    |
| [create17](#create17) | Insert vectors |
| [create18](#create18) | Query vectors  |
| [create19](#create19) | Search vectors |
| [create20](#create20) | Upsert vectors |

## All Methods

### **create**

Create a new collection

- HTTP Method: POST
- Endpoint: /collections/create

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create2**

Describe a collection

- HTTP Method: POST
- Endpoint: /collections/describe

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create2();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create3**

Drop a collection

- HTTP Method: POST
- Endpoint: /collections/drop

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create3();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create4**

Get the load state of a collection

- HTTP Method: POST
- Endpoint: /collections/get_load_state

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create4();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create5**

Get the stats of a collection

- HTTP Method: POST
- Endpoint: /collections/get_stats

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create5();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create6**

Check if a collection exists

- HTTP Method: POST
- Endpoint: /collections/has

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create6();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create7**

List all collections

- HTTP Method: POST
- Endpoint: /collections/list

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create7();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create8**

Load a collection

- HTTP Method: POST
- Endpoint: /collections/load

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create8();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create9**

Release a collection

- HTTP Method: POST
- Endpoint: /collections/release

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create9();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create10**

Rename a collection

- HTTP Method: POST
- Endpoint: /collections/rename

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.collectionsService.create10();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create11**

Create a new index

- HTTP Method: POST
- Endpoint: /indexes/create

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.indexesService.create11();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create12**

Describe an index

- HTTP Method: POST
- Endpoint: /indexes/describe

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.indexesService.create12();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create13**

Drop an index

- HTTP Method: POST
- Endpoint: /indexes/drop

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.indexesService.create13();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create14**

List all indexes

- HTTP Method: POST
- Endpoint: /indexes/list

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.indexesService.create14();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create15**

Delete vectors

- HTTP Method: POST
- Endpoint: /vectors/delete

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create15();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create16**

Get vectors

- HTTP Method: POST
- Endpoint: /vectors/get

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create16();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create17**

Insert vectors

- HTTP Method: POST
- Endpoint: /vectors/insert

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create17();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create18**

Query vectors

- HTTP Method: POST
- Endpoint: /vectors/query

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create18();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create19**

Search vectors

- HTTP Method: POST
- Endpoint: /vectors/search

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create19();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

### **create20**

Upsert vectors

- HTTP Method: POST
- Endpoint: /vectors/upsert

**Parameters**
| Name | Type| Required | Description |
| :-------- | :----------| :----------| :----------|
| input | Object | Required | Request body. |

**Return Type**

Returns a dict object.

**Example Usage Code Snippet**

```Java
package io.gandi.examples;

import io.gandi.Gandi;
import io.gandi.exceptions.ApiException;

public class Main {

  public static void main(String[] args) {
    Gandi client = new Gandi(System.getenv("GANDI_BEARER_TOKEN"));
    try {
      Object response = client.vectorsService.create20();
      System.out.println(response);
    } catch (ApiException e) {
      e.printStackTrace();
    }
  }
}

```

<!-- This file was generated by liblab | https://liblab.com/ -->
