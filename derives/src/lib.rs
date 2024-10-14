
use proc_macro::TokenStream;
use quote::quote;
use syn::{parse_macro_input, Data, DeriveInput,  Lit, Meta, NestedMeta};



#[proc_macro_derive(SiriXMLType, attributes(siri))]
pub fn xml_type_derive(input: TokenStream) -> TokenStream {
    let input = parse_macro_input!(input as DeriveInput);

    // Ensure we are working with an enum
    let name = input.ident;
    let data = match input.data {
        Data::Enum(ref data) => data,
        _ => panic!("SiriXMLType can only be derived for enums."),
    };

    let mut envelopes = vec![];

    // Iterate through each variant
    for variant in &data.variants {
        let variant_name = &variant.ident;
        let fields = match &variant.fields {
            syn::Fields::Unnamed(fields) => &fields.unnamed,
            _ => panic!("SiriXMLType can only be derived for enums with named fields."),
        };

        if fields.len() != 1 {
            panic!("SiriXMLType can only be derived for enums with a single field.");
        }
        
        for attr in &variant.attrs {
            if attr.path.is_ident("siri") {
                if let Ok(Meta::List(meta_list)) = attr.parse_meta() {
                    for nested_meta in meta_list.nested {
                        if let NestedMeta::Meta(Meta::NameValue(meta)) = nested_meta {
                            if meta.path.is_ident("file") {
                                if let Lit::Str(lit_str) = meta.lit {
                                    let file = lit_str.value();
                                    envelopes.push(quote! {
                                        #name::#variant_name(Envelope::from_file(#file).unwrap())
                                    });
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    // Implement the generated function
    let expanded = quote! {
        impl #name {
            pub fn get_profiles() -> Vec<#name> {
                vec![ #( #envelopes ),* ]
            }
        }
    };

    TokenStream::from(expanded)
}
